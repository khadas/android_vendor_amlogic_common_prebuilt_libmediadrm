ifneq ($(wildcard vendor/amlogic/common/prebuilt/libmediadrm/jcas),)
JCAS_SOURCE_PATH := $(BOARD_AML_VENDOR_PATH)/JCasLib
JCAS_PREBUILT_PATH := $(BOARD_AML_VENDOR_PATH)/prebuilt/libmediadrm/jcas

# PRODUCT_PACKAGES defines the list of installed packages.
PRODUCT_PACKAGES += libjcas_jni \
                    libjcas_native_client

PRODUCT_PACKAGES += droidlogic-jcas-builder

PRODUCT_COPY_FILES += $(JCAS_PREBUILT_PATH)/jcas_configs.json:$(TARGET_COPY_OUT_VENDOR)/etc/jcas_configs.json
endif
