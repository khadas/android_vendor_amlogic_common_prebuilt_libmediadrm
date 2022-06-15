LOCAL_PATH:= $(call my-dir)
ifeq ($(BOARD_OMX_WITH_OPTEE_TVP),true)

ifneq ($(filter arm arm64, $(TARGET_ARCH)),)

#####################################################################
# libsecmem.so
include $(CLEAR_VARS)
LOCAL_MODULE := libsecmem
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0 SPDX-license-identifier-BSD SPDX-license-identifier-LGPL legacy_by_exception_only
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted
LOCAL_PROPRIETARY_MODULE := true
LOCAL_MODULE_PATH := $(TARGET_OUT_VENDOR)/lib
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_SUFFIX := .so
LOCAL_SRC_FILES := $(LOCAL_MODULE)$(LOCAL_MODULE_SUFFIX)
LOCAL_EXPORT_C_INCLUDE_DIRS := $(LOCAL_PATH)/include
LOCAL_PROPRIETARY_MODULE := true
LOCAL_STRIP_MODULE := false
LOCAL_32_BIT_ONLY := true
LOCAL_MODULE_TAGS := optional
LOCAL_SHARED_LIBRARIES := libamavutils libc++ libc libcrypto libcutils libdl liblog libm libmediahal_resman libteec libutils libz
include $(BUILD_PREBUILT)

# libsecmem_sys.so
include $(CLEAR_VARS)
LOCAL_MODULE := libsecmem_sys
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_SUFFIX := .so
LOCAL_SRC_FILES := $(LOCAL_MODULE)$(LOCAL_MODULE_SUFFIX)
LOCAL_EXPORT_C_INCLUDE_DIRS := $(LOCAL_PATH)/include
LOCAL_STRIP_MODULE := false
LOCAL_32_BIT_ONLY := true
LOCAL_MODULE_TAGS := optional
LOCAL_CHECK_ELF_FILES := false
ifeq (1, $(strip $(shell expr $(PLATFORM_SDK_VERSION) \>= 30)))
LOCAL_SYSTEM_EXT_MODULE := true
endif
LOCAL_SHARED_LIBRARIES := libamavutils_sys libcrypto libcutils liblog libteec_sys libutils libz
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0 SPDX-license-identifier-FTL SPDX-license-identifier-GPL SPDX-license-identifier-LGPL-2.1 SPDX-license-identifier-MIT legacy_by_exception_only legacy_notice
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted
LOCAL_NOTICE_FILE := $(LOCAL_PATH)/../LICENSE
include $(BUILD_PREBUILT)

ifneq ($(USE_PRESIGNED_TA),true)
include $(CLEAR_VARS)
TA_UUID := 2c1a33c0-44cc-11e5-bc3b-0002a5d5c51b
TA_SUFFIX := .ta

ifeq ($(PLATFORM_TDK_VERSION), 38)
	PLATFORM_TDK_PATH := $(BOARD_AML_VENDOR_PATH)/tdk_v3
	ifeq ($(BOARD_AML_SOC_TYPE),)
		LOCAL_TA := ta/v3/signed/$(TA_UUID)$(TA_SUFFIX)
	else
		LOCAL_TA := ta/v3/dev/$(BOARD_AML_SOC_TYPE)/$(TA_UUID)$(TA_SUFFIX)
	endif
else
	PLATFORM_TDK_PATH := $(BOARD_AML_VENDOR_PATH)/tdk
	LOCAL_TA := ta/v2/signed/$(TA_UUID)$(TA_SUFFIX)
endif

LOCAL_SRC_FILES := $(LOCAL_TA)
LOCAL_MODULE := $(TA_UUID)
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0 SPDX-license-identifier-BSD SPDX-license-identifier-LGPL legacy_by_exception_only
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted
LOCAL_MODULE_SUFFIX := $(TA_SUFFIX)
LOCAL_STRIP_MODULE := false
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_PATH := $(TARGET_OUT_VENDOR)/lib/teetz
include $(BUILD_PREBUILT)
endif  # USE_PRESIGNED_TA != true

endif # TARGET_ARCH == arm
endif
