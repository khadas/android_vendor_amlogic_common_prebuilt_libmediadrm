/*
 * Copyright (c) 2019 Amlogic, Inc. All rights reserved.
 *
 * This source code is subject to the terms and conditions defined in the
 * file 'LICENSE' which is part of this source code package.
 *
 * Description:
 */

#ifndef AMCAS_IPTV_H
#define AMCAS_IPTV_H

#include "amCasBase.h"
#include "wvcas_types.h"

class  AmCasIPTV:public AmCasBase {

public:
    AmCasIPTV();
    ~AmCasIPTV();
    AmCasCode_t setPrivateData(void *iDate, int iSize);
    AmCasCode_t provision();
    AmCasCode_t openSession(uint8_t *sessionId);
    AmCasCode_t closeSession(uint8_t *sessionId);
    AmCasCode_t setPids(int vPid, int aPid);
    AmCasCode_t processEcm(int isSection, uint8_t *pBuffer, int iBufferLength);
    AmCasCode_t processEcm(int isSection, int iPid, uint8_t *pBuffer, int iBufferLength);//use for aml mp multi-channel cas
    AmCasCode_t processEcm(int isSection, int isVideoEcm, int vEcmPid, int aEcmPid, unsigned char *pBuffer, int iBufferLength);
    AmCasCode_t processEmm(int isSection, int iPid, uint8_t *pBuffer, int iBufferLength);
    AmCasCode_t setCasInstanceId(int casInstanceId);
    static bool isSystemIdSupported(int CA_system_id);

    #ifdef SUPPORT_CAS_PVR
    AmCasCode_t initMediaCasPvr(bool playback);
    AmCasCode_t setPvrDesType(uint32_t type);
    AmCasCode_t getPvrReScrambleKey(int32_t casSessionId,uint8_t* outKey,size_t* keyLength);
    AmCasCode_t getPvrScrambleKey(uint8_t* outKey,size_t* keyLength);
    AmCasCode_t addTrackPid(uint32_t pid);
    AmCasCode_t removeTrackPid(uint32_t pid);
    AmCasCode_t openPvrScramble(int dschandle,uint8_t* pvrKey,size_t keyLength,bool oddEvenKey);
    AmCasCode_t closeMediaCasPvr();
    #endif
    int getCasInstanceId();
    int             mCasObjIdx;//use for aml mp multi-channel cas

private:
    int             mCasInstanceId;
    unsigned int    mCaSystemId;
    unsigned int    *mEcmPid;
    unsigned int    *mEmmPid;
    char            *mLaUrl; //cas server url
    char            *mPort;//cas server port
    char            *mLicStore;//cas license store
    int             mvPid;
    int             maPid;
    CasStreamInfo   *mCasStreamInfo;
    void            *mAmcasHandle;
    #ifdef SUPPORT_CAS_PVR
    void            *mPvrCasHandle;
    #endif
};

#endif
