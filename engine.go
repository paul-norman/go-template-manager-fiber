package templateManagerFiber

import (
	"io"

	FB "github.com/gofiber/fiber/v2"
	TM "github.com/paul-norman/go-template-manager"
)

// Comply with the Fiber interface
type Engine struct {
	TM.TemplateManager
}

// Create a new instance of the View Engine
func Init(directory string, extension string) *Engine {
	return &Engine{*TM.Init(directory, extension)}
}

/*
func NewFileSystem(fs http.FileSystem, extension string) *TM.TemplateManager {
	return &Engine{
		TM.Init(directory, extension)
	}
}
*/

// Adds a custom function for use in all templates within the instance of `TemplateManager` (Fiber Naming).
func (e *Engine) AddFunc(name string, function any) *Engine {
	e.TemplateManager.AddFunction(name, function)

	return e
}

// Adds multiple custom functions for use in all templates within the instance of `TemplateManager` (Fiber Naming).
// Function names are the map keys.
func (e *Engine) AddFuncMap(functions map[string]any) *Engine {
	e.TemplateManager.AddFunctions(functions)

	return e
}

// Sets the delimiters used by `text/template` (Default: "{{" and "}}") (Fiber Naming).
func (e *Engine) Delims(left, right string) *Engine {
	e.TemplateManager.Delimiters(left, right)

	return e
}

// Triggers scanning of files and bundling of all templates
func (e *Engine) Load() error {
	return e.TemplateManager.Parse()
}

// Executes a single template (`name`)
func (e *Engine) Render(out io.Writer, template string, binding any, layout ...string) error {
	return e.TemplateManager.Render(template, parseData(binding), out)
}

// Converts generic variables to TM variables
func parseData(binding any) TM.Params {
	if binding == nil {
		return TM.Params{}
	}

	if old, ok := binding.(TM.Params); ok {
		return old
	}

	if old, ok := binding.(map[string]any); ok {
		return TM.Params(old)
	}

	if old, ok := binding.(FB.Map); ok {
		return TM.Params(old)
	}

	return TM.Params{}
}