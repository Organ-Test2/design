package parser

import (
	. "../languages/design"
	"fmt"
)

func NewDesignAppListener() *DesignAppListener {
	return &DesignAppListener{}
}

type DesignAppListener struct {
	BaseDesignListener
}


func (s *DesignAppListener) EnterConfigDeclaration(ctx *ConfigDeclarationContext) {
	fmt.Println(ctx.ConfigKey().GetText(), ctx.ConfigValue().GetText())
}

func (s *DesignAppListener) EnterComponentName(ctx *ComponentNameContext) {
	//fmt.Println(ctx)
}

func (s *DesignAppListener) EnterLayoutDecalaration(ctx *LayoutDecalarationContext) {
	//fmt.Println(ctx.GetText())
}

func (s *DesignAppListener) EnterLayoutRow(ctx *LayoutRowContext) {
	//fmt.Println(ctx.GetText())
}

func (s *DesignAppListener) EnterComponentUseDeclaration(ctx *ComponentUseDeclarationContext) {
	//fmt.Println(ctx.GetText())
}