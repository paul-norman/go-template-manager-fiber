package templateManagerFiber

import (
	"fmt"
	"io"

	FB "github.com/gofiber/fiber/v2"
	TM "github.com/paul-norman/go-template-manager"
)

// Comply with the Fiber interface
type Engine struct {
	tm TM.TemplateManager
}

// Create a new instance of the View Engine
func New(directory string, extension string) *Engine {
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
	e.tm.AddFunction(name, function)

	return e
}

// Adds a custom function for use in all templates within the instance of `TemplateManager`.
func (e *Engine) AddFunction(name string, function any) *Engine {
	e.tm.AddFunction(name, function)

	return e
}

// Adds multiple custom functions for use in all templates within the instance of `TemplateManager` (Fiber Naming).
// Function names are the map keys.
func (e *Engine) AddFuncMap(functions map[string]any) *Engine {
	e.tm.AddFunctions(functions)

	return e
}

// Adds multiple custom functions for use in all templates within the instance of `TemplateManager`.
// Function names are the map keys.
func (e *Engine) AddFunctions(functions map[string]any) *Engine {
	e.tm.AddFunctions(functions)

	return e
}

// Adds a single variable (`name`) with value `value` that will always be available in the `templateName` template
func (e *Engine) AddParam(templateName string, name string, value any) *Engine {
	e.tm.AddParam(templateName, name, value)

	return e
}

// Adds several variables (`params`) that will always be available in the `templateName` template
func (e *Engine) AddParams(templateName string, params TM.Params) *Engine {
	e.tm.AddParams(templateName, params)

	return e
}

// Sets the delimiters used by `text/template` (Default: "{{" and "}}") (Fiber Naming).
func (e *Engine) Delims(left, right string) *Engine {
	e.tm.Delimiters(left, right)

	return e
}

// Sets the delimiters used by `text/template` (Default: "{{" and "}}")
func (e *Engine) Delimiters(left string, right string) *Engine {
	e.tm.Delimiters(left, right)

	return e
}

// Enable debugging of the template build process
func (e *Engine) Debug(enabled bool) *Engine {
	e.tm.Debug(enabled)

	return e
}

// Excludes multiple directories from the build scanning process (which only wants entry files).
// This does not prevent files in these directories from being included via `template`.
// Typically, directories containing base layouts and partials should be excluded.
func (e *Engine) ExcludeDirectories(directories []string) *Engine {
	e.tm.ExcludeDirectories(directories)

	return e
}

// Exclude a directory from the build scanning process (which only wants entry files).
// This does not prevent files in this directory from being included via `template`.
// Typically, directories containing base layouts and partials should be excluded.
func (e *Engine) ExcludeDirectory(directory string) *Engine {
	e.tm.ExcludeDirectory(directory)

	return e
}

// Not required
func (e *Engine) Layout(key string) *Engine {
	return e
}

// Triggers scanning of files and bundling of all templates
func (e *Engine) Load() error {
	return e.tm.Parse()
}

// Not required
func (e *Engine) Parse() error {
	return fmt.Errorf("Parse() is deprecated, please use Load() instead")
}

// Removes a directory that was previously excluded to allow it to feature in the build scanning process (which only wants entry files).
func (e *Engine) RemoveExcludedDirectory(directory string) *Engine {
	e.tm.RemoveExcludedDirectory(directory)

	return e
}

// Removes all functions currently assigned to the instance of `TemplateManager`.
// Useful if you do not want the default functions included
func (e *Engine) RemoveAllFunctions() *Engine {
	e.tm.RemoveAllFunctions()

	return e
}

// Enable re-rebuilding of the template bundle upon every page load (for development)
func (e *Engine) Reload(enabled bool) *Engine {
	e.tm.Reload(enabled)

	return e
}

// Executes a single template (`name`)
func (e *Engine) Render(out io.Writer, template string, binding any, layout ...string) error {
	return e.tm.Render(template, convertFiberToTM(binding), out)
}

// Converts Fiber variables to TM variables
func convertFiberToTM(binding any) TM.Params {
	old := binding.(FB.Map)
	new := TM.Params{}
	for key, value := range old {
		new[key] = value
	}

	return new
}