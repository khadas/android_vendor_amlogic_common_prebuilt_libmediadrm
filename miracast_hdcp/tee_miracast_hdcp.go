package tee_miracast_hdcp

import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "github.com/google/blueprint/proptools"
    "android/soong/vendor/amlogic/common/prebuilt/libmediadrm/common"
)

func init() {
    android.RegisterModuleType("libstagefright_go_defaults", libstagefright_DefaultsFactory)
}

func libstagefright_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "libstagefright_hdcp.so",
            SrcPath: "vendor/amlogic/common/miracast-hdcp2",
            System: false,
            MultiArch: true,
        }
        props.Enabled = proptools.BoolPtr(true)
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("libstagefright_hdcp p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("libstagefright_hdcp p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}
