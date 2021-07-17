#define LOG_TAG "NetflixHidlService"

#include <android-base/logging.h>
#include <binder/ProcessState.h>
#include <hidl/HidlTransportSupport.h>

#include "NetflixFactory.h"

using vendor::hardware::drm::V1_4::netflix::NFCryptoFactory;
using vendor::hardware::drm::V1_4::netflix::NFDrmFactory;

int main(int /* argc */, char** /* argv */) {
    sp<IDrmFactory> drmFactory = new NFDrmFactory;
    sp<ICryptoFactory> cryptoFactory = new NFCryptoFactory;

    configureRpcThreadpool(8, true /* callerWillJoin */);

    android::hardware::setRequestingSid(drmFactory, true);
    android::hardware::setRequestingSid(cryptoFactory, true);

    // Setup hwbinder service
    CHECK_EQ(drmFactory->registerAsService("netflix"), android::NO_ERROR)
        << "Failed to register Netflix Factory HAL";
    CHECK_EQ(cryptoFactory->registerAsService("netflix"), android::NO_ERROR)
        << "Failed to register Netflix Crypto  HAL";

    joinRpcThreadpool();
}
