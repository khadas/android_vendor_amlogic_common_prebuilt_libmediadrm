#ifndef NETFLIX_FACTORY_H_
#define NETFLIX_FACTORY_H_

#include "HidlTypes.h"

using namespace android;

namespace vendor {
namespace hardware {
namespace drm {
namespace V1_4 {
namespace netflix {

struct NFCryptoFactory : public ::drm::V1_4::ICryptoFactory {
 public:
  NFCryptoFactory();
  virtual ~NFCryptoFactory();

  Return<bool> isCryptoSchemeSupported(const hidl_array<uint8_t, 16>& uuid)
      override;

  Return<void> createPlugin(
      const hidl_array<uint8_t, 16>& uuid,
      const hidl_vec<uint8_t>& initData,
      createPlugin_cb _hidl_cb) override;

 private:
    void * mCryptoFactory;
};

struct NFDrmFactory : public ::drm::V1_4::IDrmFactory {
  NFDrmFactory();
  virtual ~NFDrmFactory();

  Return<bool> isCryptoSchemeSupported(const hidl_array<uint8_t, 16>& uuid)
      override;

  Return<bool> isCryptoSchemeSupported_1_2(const hidl_array<uint8_t, 16>& uuid,
                                           const hidl_string& mimeType,
                                           SecurityLevel level)
      override;

  Return<bool> isContentTypeSupported(const hidl_string &mimeType)
      override;

  Return<void> createPlugin(
      const hidl_array<uint8_t, 16>& uuid,
      const hidl_string& appPackageName,
      createPlugin_cb _hidl_cb) override;

  Return<void> getSupportedCryptoSchemes(
      getSupportedCryptoSchemes_cb _hidl_cb) override;

 private:
    void * mPluginFactory;
};

extern "C" IDrmFactory* HIDL_FETCH_IDrmFactory(const char* name);

}  // namespace vendor
}  // namespace V1_4
}  // namespace drm
}  // namespace hardware
}  // namespace netflix

#endif // NETFLIX_FACTORY_H_