package droidlogic_jcas
import (
    "fmt"
    "android/soong/android"
    "android/soong/cc"
    "github.com/google/blueprint/proptools"
)

func init() {
    android.RegisterModuleType("droidlogic_jcas_jni_go_defaults", jcas_jni_go_DefaultsFactory)
}

func jcas_jni_go_DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, func(ctx android.LoadHookContext) {
        type props struct {
            Enabled *bool
        }
        p := &props{}

        jcasSrcPath := "vendor/amlogic/common/JCasLib"
        if android.ExistentPathForSource(ctx, jcasSrcPath).Valid() == true {
            p.Enabled = proptools.BoolPtr(false)
            fmt.Printf("jcasSrcPath:%s exist, use jcas source to build\n", jcasSrcPath)
        } else {
            p.Enabled = proptools.BoolPtr(true)
            fmt.Printf("jcasSrcPath:%s not exist, use jcas prebuilt to build\n", jcasSrcPath)
        }
        ctx.AppendProperties(p)
    })
    return module
}