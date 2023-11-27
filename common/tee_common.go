package common

import (
    "fmt"
    "strings"
    "android/soong/android"
    "github.com/google/blueprint/proptools"
)
/* output props*/
type Props struct {
    Enabled *bool
    system_ext_specific *bool
    Arch struct {
        Arm struct {
            Srcs []string
        }
        Arm64 struct {
            Srcs []string
        }
    }
}

/*input Info*/
type Information struct {
    FileString string
    SrcPath string
    System bool
    MultiArch bool
}

func stringToSlice(input string) []string {
    result := strings.Split(input, " ")
    return result
}

func SetProps(ctx android.LoadHookContext, props *Props, info Information) {
    vconfig := ctx.Config().VendorConfig("amlogic_vendorconfig")
    tdk_ver := vconfig.String("tdk_version")
    if android.ExistentPathForSource(ctx, info.SrcPath).Valid() == true {
        props.Enabled = proptools.BoolPtr(false)
        fmt.Printf("%s exist, use  source to build\n", info.SrcPath)
    }
    if (info.System) {
        sdkVersion := ctx.Config().PlatformSdkVersion()
        if (sdkVersion.FinalOrFutureInt() > 30) {
            props.system_ext_specific = proptools.BoolPtr(true)
        }
    }
    if tdk_ver == "TDK24" || tdk_ver == "TDK38" {
        tdk_ver = "v3.8.0"
    } else if tdk_ver == "TDK318" {
        tdk_ver = "v3.18.0"
    } else {
        tdk_ver = "v3.8.0"
        fmt.Printf("tek_ver not exist, use TDK38\n")
    }
    props.Arch.Arm.Srcs = stringToSlice(fmt.Sprintf("%s/%s/%s", tdk_ver,
                                                           "arm",info.FileString))
    if (info.MultiArch) {
        props.Arch.Arm64.Srcs = stringToSlice(fmt.Sprintf("%s/%s/%s", tdk_ver,
                                                                 "arm64", info.FileString))
    }
}