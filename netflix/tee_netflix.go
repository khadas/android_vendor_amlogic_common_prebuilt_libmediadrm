package tee_netflix
import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "android/soong/vendor/amlogic/common/prebuilt/libmediadrm/common"
)

func init() {
    android.RegisterModuleType("libtee_hal_go_defaults", libtee_hal_DefaultsFactory)
}

func libtee_hal_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "libtee-hal.so",
            SrcPath: "vendor/amlogic/common/netflix",
            System: false,
            MultiArch: true,
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("libtee-hal p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("libtee-hal p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}

