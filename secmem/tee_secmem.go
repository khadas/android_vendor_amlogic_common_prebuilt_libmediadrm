package tee_secmem
import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "github.com/google/blueprint/proptools"
    "android/soong/vendor/amlogic/common/prebuilt/libmediadrm/common"
)

func init() {
    android.RegisterModuleType("libsecmem_go_defaults",   libsecmem_DefaultsFactory)
    android.RegisterModuleType("libsecmem_sys_go_defaults", libsecmem_sys_DefaultsFactory)
}

func libsecmem_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "libsecmem.so",
            SrcPath: "vendor/amlogic/common/libsecmem",
            System: false,
            MultiArch: true,
        }
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        props.Enabled = proptools.BoolPtr(true)
        if vconfig.Bool("omx_with_optee_tvp") == false {
            // fmt.Println("secmem: omx_with_optee_tvp is false, not include prebuilt")
            props.Enabled = proptools.BoolPtr(false)
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("libsecmem p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("libsecmem p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}

func libsecmem_sys_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        var props common.Props
        info := common.Information{
            FileString: "libsecmem_sys.so",
            SrcPath: "vendor/amlogic/common/libsecmem",
            System: true,
            MultiArch: true,
        }
        vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
        props.Enabled = proptools.BoolPtr(true)
        if vconfig.Bool("omx_with_optee_tvp") == false {
            // fmt.Println("secmem: omx_with_optee_tvp is false, not include prebuilt")
            props.Enabled = proptools.BoolPtr(false)
        }
        common.SetProps(ctx, &props, info)
        p := &props
        fmt.Println("libsecmem_sys p.Arch.Arm.Srcs", p.Arch.Arm.Srcs)
        fmt.Println("libsecmem_sys p.Arch.Arm64.Srcs", p.Arch.Arm64.Srcs)
        ctx.AppendProperties(p)
    })
    return module
}
