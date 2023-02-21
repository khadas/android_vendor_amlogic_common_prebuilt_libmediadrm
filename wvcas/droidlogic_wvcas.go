package droidlogic_wvcas
import (
    //"fmt"
    "android/soong/android"
    "android/soong/cc"
    "github.com/google/blueprint/proptools"
)

func init() {
    android.RegisterModuleType("droidlogic_wvcas_go_defaults",      wvcas_go_DefaultsFactory)
}

func wvcas_go_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        type props struct {
            Enabled *bool
        }
        p := &props{}

        wvcasSrcPath := "vendor/amlogic/common/widevinecas"
        if android.ExistentPathForSource(ctx, wvcasSrcPath).Valid() == true {
            p.Enabled = proptools.BoolPtr(false)
            //fmt.Printf("wvcasSrcPath:%s exist, use wvcas source to build\n", wvcasSrcPath)
        } else {
            //fmt.Printf("wvcasSrcPath:%s not exist, use wvcas prebuilt to build\n", wvcasSrcPath)

            vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
            if vconfig.String("widevine_oemcrypto_level") != "1" {
                //fmt.Println("wvcas: widevine_oemcrypto_level != 1, not include wvcas prebuilt")
                p.Enabled = proptools.BoolPtr(false)
            } else {
                //fmt.Println("wvcas: widevine_oemcrypto_level == 1 can build wvcas prebuilt")
            }
        }
        ctx.AppendProperties(p)
    })
    return module
}