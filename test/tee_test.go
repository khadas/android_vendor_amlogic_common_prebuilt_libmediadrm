package tee_test

import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "android/soong/vendor/amlogic/common/prebuilt/libmediadrm/common"
)

func init() {
    android.RegisterModuleType("tee_secmem_test_go_defaults",
                                    tee_secmem_test_DefaultsFactory)
    android.RegisterModuleType("tee_wvcas_iptv_test_go_defaults",
                                    tee_wvcas_iptv_test_DefaultsFactory)
    android.RegisterModuleType("tee_wvcas_iptv_test_sys_go_defaults",
                                    tee_wvcas_iptv_test_sys_DefaultsFactory)
}

func tee_secmem_test_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "secmem_test",
            SrcPath: "vendor/amlogic/common/libsecmem",
            System: false,
            MultiArch: true,
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("secmem_test p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("secmem_test p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}

func tee_wvcas_iptv_test_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "wvcas_iptv_test",
            SrcPath: "vendor/amlogic/common/widevinecas",
            System: false,
            MultiArch: false,
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("tee_wvcas_iptv_test p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}

func tee_wvcas_iptv_test_sys_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "wvcas_iptv_test_sys",
            SrcPath: "vendor/amlogic/common/widevinecas",
            System: false,
            MultiArch: false,
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("tee_wvcas_iptv_test_sys p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}
