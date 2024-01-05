LOCAL_PATH:= $(call my-dir)

ifneq ($(USE_PRESIGNED_TA),true)
include $(CLEAR_VARS)
TA_UUID := 1077efec-c0b2-4d02-ace3-3c1e52e2fb4b
TA_SUFFIX := .ta

ifneq ($(PLATFORM_TDK_VERSION), 24)
    ifeq ($(BOARD_AML_SOC_TYPE),)
        LOCAL_TA := ta/v3/signed/$(TA_UUID)$(TA_SUFFIX)
    else
        LOCAL_TA := ta/v3/dev/$(BOARD_AML_SOC_TYPE)/$(TA_UUID)$(TA_SUFFIX)
    endif
else
    LOCAL_TA :=  ta/v2/signed/$(TA_UUID)$(TA_SUFFIX)
endif

LOCAL_SRC_FILES := $(LOCAL_TA)
LOCAL_MODULE := $(TA_UUID)
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0 legacy_proprietary
LOCAL_LICENSE_CONDITIONS := notice proprietary by_exception_only
LOCAL_MODULE_SUFFIX := $(TA_SUFFIX)
LOCAL_STRIP_MODULE := false
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_PATH := $(TARGET_OUT_VENDOR)/lib/teetz
LOCAL_32_BIT_ONLY := true
include $(BUILD_PREBUILT)
endif  # USE_PRESIGNED_TA != true
