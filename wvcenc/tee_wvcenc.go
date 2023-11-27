package tee_wvcenc
import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "github.com/google/blueprint/proptools"
    "android/soong/vendor/amlogic/common/prebuilt/libmediadrm/common"
)

func init() {
    android.RegisterModuleType("liboemcrypto_go_defaults", liboemcrypto_DefaultsFactory)
}

func liboemcrypto_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "liboemcrypto.so",
            SrcPath: "vendor/amlogic/common/widevine",
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
        fmt.Println("liboemcrypto p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("liboemcrypto p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}
