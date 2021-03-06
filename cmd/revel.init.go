package cmd

import (
	"log"
	"path"

	"github.com/liujianping/scaffold/symbol"
)

func revel_init(project string, template_dir string, force bool) error {
	root_dir, err := gopath_src_dir()
	if err != nil {
		return err
	}
	project_dir := path.Join(root_dir, project)

	data := map[string]interface{}{
		"project": project,
	}

	if err := symbol.CopyDir(path.Join(template_dir, "revel", "public"),
		path.Join(project_dir, "public")); err != nil {
		return err
	}
	log.Println("copy directory =>", path.Join(project_dir, "public"))

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "models", "model.go.t"),
		path.Join(project_dir, "app", "models", "model.go"), data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "models", "model.validator.go.t"),
		path.Join(project_dir, "app", "models", "model.validator.go"), data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "controllers", "controller.widget.go.t"),
		path.Join(project_dir, "app", "controllers", "controller.widget.go"), data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "controllers", "controller.db.go.t"),
		path.Join(project_dir, "app", "controllers", "controller.db.go"), data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "controllers", "controller.fns.go.t"),
		path.Join(project_dir, "app", "controllers", "controller.fns.go"), data, force); err != nil {
		return err
	}

	return nil
}
