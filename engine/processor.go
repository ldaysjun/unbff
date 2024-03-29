package engine

import (
	"bytes"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/source"
	"github.com/ldaysjun/unbff/config"
	"github.com/ldaysjun/unbff/dao"
	"github.com/ldaysjun/unbff/model"
	"html/template"
	"log"
	"xorm.io/core"
	"xorm.io/xorm"
)

type processor struct {
	db          *xorm.Engine
	objects     map[string]map[string]*graphql.Object
	appDataList map[string][]*model.SDLMetadata
}

func newProcessor(conf *config.Config) *processor {
	p := &processor{}
	p.objects = make(map[string]map[string]*graphql.Object)
	p.db = newProcessorDB(conf)
	return p
}

func (p *processor) generateRootFields() (map[string]graphql.Fields, error) {
	sdlList, err := dao.SDLMetadataList(p.db)
	if err != nil {
		return nil, err
	}
	fields := make(map[string]graphql.Fields)
	p.appDataList = groupSDL(sdlList)
	for app, metadataList := range p.appDataList {
		fields[app] = p.appFields(aggregateDocuments(metadataList), app)
	}
	return fields, nil
}

func (p *processor) appFields(document *ast.Document, app string) graphql.Fields {
	// interface node, deal with last
	queryNodes := make([]*ast.ObjectDefinition, 0)
	for _, definition := range document.Definitions {
		node, ok := definition.(*ast.ObjectDefinition)
		if !ok {
			log.Println("node is not ast.ObjectDefinition")
			continue
		}
		name := node.Name.Value
		// filter interface node
		if isInterface(name) {
			queryNodes = append(queryNodes, node)
			continue
		}
		// generate graphql object, insert to objectSet
		if _, ok := p.objects[app][name]; !ok {
			obj := p.generateGraphQLObject(app, node)
			if _, ok := p.objects[app]; !ok {
				p.objects[app] = make(map[string]*graphql.Object)
			}
			p.objects[app][obj.Name()] = obj
		}
	}

	// last generated interface field
	fields := p.generateQueryFields(app, queryNodes)
	return fields
}

func (p *processor) generateQueryFields(app string, nodes []*ast.ObjectDefinition) graphql.Fields {
	fieldDefinitions := make([]*ast.FieldDefinition, 0)
	for _, n := range nodes {
		fieldDefinitions = append(fieldDefinitions, n.Fields...)
	}
	fields := make(graphql.Fields)
	for _, field := range fieldDefinitions {
		fields[field.Name.Value] = &graphql.Field{
			Type: p.generateGraphQLType(app, field.Type),
			Args: p.generateGraphQLArgs(app, field.Arguments),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return p.resolve(params, field.Directives)
			},
		}
	}
	return fields
}

// resolve : design universal resolve
func (p *processor) resolve(params graphql.ResolveParams,
	directives []*ast.Directive) (interface{}, error) {
	var n node
	for _, directive := range directives {
		if directive.Name.Value == "rest" {
			decorate := httpNodeDecorate()
			hn := &httpNode{
				args: params.Args,
			}
			for _, argument := range directive.Arguments {
				value := argument.Name.Value
				if function, ok := decorate[value]; ok {
					v := fmt.Sprintf("%v", argument.Value.GetValue())
					function(hn, v, params.Args)
				}
			}
			n = hn
		}
		result, err := n.do()
		if err != nil {
			return map[string]interface{}{}, err
		}
		if result == nil {
			return map[string]interface{}{}, nil
		}
		return result, nil
	}
	return nil, nil
}

// httpNodeDecorate http node decorate
func httpNodeDecorate() map[string]func(n *httpNode, v string, args map[string]interface{}) {
	return map[string]func(n *httpNode, v string, args map[string]interface{}){
		"method": func(n *httpNode, v string, args map[string]interface{}) {
			n.method = v
		},
		"contenttype": func(n *httpNode, v string, args map[string]interface{}) {
			n.contentType = v
		},
		"endpoint": func(n *httpNode, v string, args map[string]interface{}) {
			n.endpoint = v
		},
		"root": func(n *httpNode, v string, args map[string]interface{}) {
			n.root = v
		},
		"headers": func(n *httpNode, v string, args map[string]interface{}) {
			n.headers = v
		},
		"body": func(n *httpNode, v string, args map[string]interface{}) {
			tmpl, err := template.New("tmpl").Parse(v)
			if err != nil {
				fmt.Println("create template failed, err:", err)
				return
			}
			var tmplBytes bytes.Buffer
			if err := tmpl.Execute(&tmplBytes, args); err != nil {
				fmt.Println(err)
				return
			}
			n.postBody = tmplBytes.String()
		},
	}
}

func (p *processor) generateGraphQLArgs(app string,
	arguments []*ast.InputValueDefinition) graphql.FieldConfigArgument {
	args := make(graphql.FieldConfigArgument)
	for _, argument := range arguments {
		args[argument.Name.Value] = &graphql.ArgumentConfig{
			Type: p.generateGraphQLType(app, argument.Type),
		}
	}
	return args
}

func (p *processor) dependGraphQLObject(app, objName string) *graphql.Object {
	document := aggregateDocuments(p.appDataList[app])
	for _, definition := range document.Definitions {
		node, ok := definition.(*ast.ObjectDefinition)
		if !ok {
			log.Println("node is not ast.ObjectDefinition")
			continue
		}
		if node.Name.Value != objName {
			continue
		}
		if isInterface(node.Name.Value) {
			continue
		}
		obj := p.generateGraphQLObject(app, node)
		p.objects[app][obj.Name()] = obj
	}
	return nil
}

func (p *processor) generateGraphQLObject(app string, definition *ast.ObjectDefinition) *graphql.Object {
	if obj, ok := p.objects[app][definition.Name.Value]; ok {
		return obj
	}
	fields := make(graphql.Fields)
	for _, f := range definition.Fields {
		field := p.generateGraphQLFiled(app, f)
		fields[field.Name] = field
	}
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name:   definition.Name.Value,
		Fields: fields,
	})
	return obj
}

func (p *processor) generateGraphQLFiled(app string, definition *ast.FieldDefinition) *graphql.Field {
	if definition.Kind != fieldDefinitionKind {
		return nil
	}
	field := &graphql.Field{
		Name: definition.Name.Value,
		Type: p.generateGraphQLType(app, definition.Type),
	}
	return field
}

func (p *processor) generateGraphQLType(app string, t ast.Type) graphql.Type {
	switch t.GetKind() {
	case ListKind:
		listType := t.(*ast.List)
		return graphql.NewList(p.generateGraphQLType(app, listType.Type))
	case NamedKind:
		namedType := t.(*ast.Named)
		v := namedType.Name.Value
		return p.namedType(app, v)
	case NonNullKind:
		nonNullType := t.(*ast.NonNull)
		return p.generateGraphQLType(app, nonNullType.Type)
	default:
		return nil
	}
}

func (p *processor) namedType(app string, v string) graphql.Type {
	baseType := map[string]graphql.Type{
		StringType:  graphql.String,
		IntType:     graphql.Int,
		FloatType:   graphql.Float,
		BooleanType: graphql.Boolean,
	}

	if t, ok := baseType[v]; ok {
		return t
	}
	if gt, ok := p.objects[app][v]; ok {
		return gt
	}
	depObj := p.dependGraphQLObject(app, v)
	return depObj
}

func newProcessorDB(conf *config.Config) *xorm.Engine {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		conf.ProcessorDB.DBUser, conf.ProcessorDB.DBPwd, conf.ProcessorDB.DBHost, conf.ProcessorDB.DBPort,
		conf.ProcessorDB.DBName, conf.ProcessorDB.DBCharset)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	engine.SetMapper(core.SameMapper{})
	engine.ShowSQL(true)
	if err := engine.Ping(); err != nil {
		panic(err)
	}
	return engine
}

func groupSDL(dataList []*model.SDLMetadata) map[string][]*model.SDLMetadata {
	temp := make(map[string][]*model.SDLMetadata)
	for _, d := range dataList {
		temp[d.App] = append(temp[d.App], d)
	}
	return temp
}

func aggregateDocuments(metadataList []*model.SDLMetadata) *ast.Document {
	document := &ast.Document{Kind: "AllDocument"}
	for _, metadata := range metadataList {
		a, err := parser.Parse(parser.ParseParams{Source: &source.Source{
			Body: []byte(metadata.SDL),
			Name: "GraphQL",
		}})
		if err != nil {
			panic(err)
		}
		document.Definitions = append(document.Definitions, a.Definitions...)
	}
	return document
}

func isInterface(name string) bool {
	if len(name) < len(query) {
		return false
	}
	return name[len(name)-len(query):] == query
}
