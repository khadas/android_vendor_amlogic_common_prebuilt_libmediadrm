LOCAL_PATH:= $(call my-dir)
ifeq ($(TARGET_BUILD_NETFLIX_MGKID), true)
#####################################################################
# libtee-hal.so
include $(CLEAR_VARS)
LOCAL_MODULE := libtee-hal
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0 SPDX-license-identifier-BSD SPDX-license-identifier-LGPL legacy_by_exception_only
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted
LOCAL_SYSTEM_EXT_MODULE := true
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_SUFFIX := .so
LOCAL_SRC_FILES_32 := arm/$(LOCAL_MODULE)$(LOCAL_MODULE_SUFFIX)
LOCAL_SRC_FILES_64 := arm64/$(LOCAL_MODULE)$(LOCAL_MODULE_SUFFIX)
LOCAL_STRIP_MODULE := false
LOCAL_MODULE_TAGS := optional
LOCAL_CHECK_ELF_FILES := false
LOCAL_SHARED_LIBRARIES := libcutils liblog libteec_sys libutils
include $(BUILD_PREBUILT)

ifneq ($(USE_PRESIGNED_TA),true)
# 00d1ca22-1764-4e35-90aa-5b8c12630764.ta
include $(CLEAR_VARS)
TA_UUID := 00d1ca22-1764-4e35-90aa-5b8c12630764
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
LOCAL_32_BIT_ONLY := true
LOCAL_LICENSE_KINDS := SPDX-license-identifier-Apache-2.0 SPDX-license-identifier-BSD SPDX-license-identifier-LGPL legacy_by_exception_only
LOCAL_LICENSE_CONDITIONS := by_exception_only notice restricted
LOCAL_MODULE_SUFFIX := $(TA_SUFFIX)
LOCAL_STRIP_MODULE := false
LOCAL_MODULE_CLASS := SHARED_LIBRARIES
LOCAL_MODULE_PATH := $(TARGET_OUT_VENDOR)/lib/teetz
include $(BUILD_PREBUILT)
endif  # USE_PRESIGNED_TA != true

# -----------------------------------------------------------------------------
# Builds android.hardware.drm@1.4-service.netflix
#
#ifneq (0, $(shell expr $(PLATFORM_SDK_VERSION) \>= 30))
#include $(CLEAR_VARS)

#LOCAL_C_INCLUDES := \
#  $(LOCAL_PATH) \
#  $(LOCAL_PATH)/include

#LOCAL_SHARED_LIBRARIES := \
#  android.hardware.drm@1.0 \
#  android.hardware.drm@1.1 \
#  android.hardware.drm@1.2 \
#  android.hardware.drm@1.3 \
#  android.hardware.drm@1.4 \
#  libbase \
#  libhidlbase \
#  liblog \
#  libutils \
#  libbinder \
#  libnetflixplugin

#LOCAL_MODULE_PATH := $(TARGET_OUT_VENDOR)/bin/hw
#LOCAL_PROPRIETARY_MODULE := true

#LOCAL_SRC_FILES := service.cpp

#SERVICE_VERSION = 1.4

#LOCAL_MODULE := android.hardware.drm@$(SERVICE_VERSION)-service.netflix
#LOCAL_INIT_RC := android.hardware.drm@$(SERVICE_VERSION)-service.netflix.rc
#LOCAL_VINTF_FRAGMENTS := manifest_android.hardware.drm@$(SERVICE_VERSION)-service.netflix.xml
#include $(BUILD_EXECUTABLE)
#endif
endif
