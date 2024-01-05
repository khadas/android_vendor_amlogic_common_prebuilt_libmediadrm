package tee_clearkey
import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "github.com/google/blueprint/proptools"
    "android/soong/vendor/amlogic/common/prebuilt/libmediadrm/common"
)

func init() {
    android.RegisterModuleType("libclearkeydrm_go_defaults",
                                      libclearkeydrm_DefaultsFactory)
    android.RegisterModuleType("tee_clearkey_go_defaults",
                                      tee_clearkey_DefaultsFactory)
}

func libclearkeydrm_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "libclearkeydrm.so",
            SrcPath: "vendor/amlogic/refenence/clearkey",
            System: false,
            MultiArch: true,
        }
        props.Enabled = proptools.BoolPtr(true)
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("libclearkeydrm p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("libclearkeydrm p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}

func tee_clearkey_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "android.hardware.drm-service.clearkey.secure",
            SrcPath: "vendor/amlogic/refenence/clearkey",
            System: false,
            MultiArch: true,
        }
        props.Enabled = proptools.BoolPtr(true)
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("drm-service.clearkey p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("drm-service.clearkey p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}