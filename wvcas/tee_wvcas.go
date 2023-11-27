package tee_wvcas
import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "github.com/google/blueprint/proptools"
    "android/soong/vendor/amlogic/common/prebuilt/libmediadrm/common"
)

func init() {
    android.RegisterModuleType("libwvmediacas_go_defaults", libwvmediacas_DefaultsFactory)
    android.RegisterModuleType("libdec_ca_wvcas_go_defaults", libdec_ca_wvcas_DefaultsFactory)
    android.RegisterModuleType("libdec_ca_wvcas_system_go_defaults", libdec_ca_wvcas_system_DefaultsFactory)
}

func libwvmediacas_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "libwvmediacas.so",
            SrcPath: "vendor/widevine/libwvmediacas",
            System: false,
            MultiArch: true,
        }
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        props.Enabled = proptools.BoolPtr(true)
        if vconfig.String("widevine_oemcrypto_level") != "1" {
            //fmt.Println("wv: widevine_oemcrypto_level != 1, not include wv prebuilt")
            props.Enabled = proptools.BoolPtr(false)
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("libwvmediacas p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("libwvmediacas p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}

func libdec_ca_wvcas_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "libdec_ca_wvcas.so",
            SrcPath: "vendor/amlogic/common/widevinecas",
            System: false,
            MultiArch: false,
        }
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        props.Enabled = proptools.BoolPtr(true)
        if vconfig.String("widevine_oemcrypto_level") != "1" {
            //fmt.Println("wv: widevine_oemcrypto_level != 1, not include wv prebuilt")
            props.Enabled = proptools.BoolPtr(false)
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("libdec_ca_wvcas p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}

func libdec_ca_wvcas_system_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "libdec_ca_wvcas.system.so",
            SrcPath: "vendor/amlogic/common/widevinecas",
            System: false,
            MultiArch: false,
        }
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        props.Enabled = proptools.BoolPtr(true)
        if vconfig.String("widevine_oemcrypto_level") != "1" {
            //fmt.Println("wv: widevine_oemcrypto_level != 1, not include wv prebuilt")
            props.Enabled = proptools.BoolPtr(false)
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("libdec_ca_wvcas.system p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}
