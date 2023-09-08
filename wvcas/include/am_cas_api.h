/*
 * Copyright (c) 2021 Amlogic, Inc. All rights reserved.
 *
 * This source code is subject to the terms and conditions defined in the
 * file 'LICENSE' which is part of this source code package.
 *
 * Description:
 */

#ifndef AMCAS_API_H
#define AMCAS_API_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#define AM_CAS_ChkArgs(expr) do { \
        if (expr) { \
            ALOGE("[%s:%d] error params!\n", __FUNCTION__, __LINE__); \
            return CAS_STATUS_ERROR; \
        } \
    } while(0);

typedef enum
{
    CAS_STATUS_OK,
    CAS_STATUS_ERROR,
    CAS_STATUS_INIT_FAILURE,
    CAS_STATUS_PARAMETER_ERROR,
    CAS_STATUS_NOT_IMPLEMENTED,
} AmCasStatus_t;

AmCasStatus_t createAmCas(void **casObj);

AmCasStatus_t setCasInstanceId(void *casObj, int casInstanceId);

int getCasInstanceId(void *casObj);

AmCasStatus_t setPrivateData(void *casObj, void *iDate, int iSize);

AmCasStatus_t provision(void *casObj);

AmCasStatus_t setPids(void *casObj, int vPid, int aPid);

AmCasStatus_t processEcm(void *casObj, int isSection, int isVideoEcm, int vEcmPid, int aEcmPid, \
                            unsigned char *pBuffer, int iBufferLength);

AmCasStatus_t processEmm(void *casObj, int isSection,
    int iPid, uint8_t *pBuffer, int iBufferLength);

AmCasStatus_t openSession(void *casObj, uint8_t *sessionId);

AmCasStatus_t decrypt(void *casObj, uint8_t *in, uint8_t *out, int size, void *ext_data);

uint8_t* getOutbuffer(void *casObj);

AmCasStatus_t closeSession(void *casObj, uint8_t *sessionId);

AmCasStatus_t releaseAll(void *casObj);

bool isSystemIdSupported(int CA_system_id);

#ifdef SUPPORT_CAS_PVR
AmCasStatus_t initMediaCasPvr(void *casObj,bool playback);
AmCasStatus_t getPvrReScrambleKey(void *casObj,int32_t casSessionId,uint8_t* outKey,size_t* keyLength);
AmCasStatus_t getPvrScrambleKey(void *casObj,uint8_t* outKey,size_t* keyLength);
AmCasStatus_t setPvrDescrambleType(void *casObj,uint32_t type);
AmCasStatus_t openPvrScramble(void *casObj,int dschandle,uint8_t* pvrKey,size_t keyLength,bool oddEvenKey);
AmCasStatus_t addTrackPid(void *casObj, uint32_t pid);
AmCasStatus_t removeTrackPid(void *casObj, uint32_t pid);
AmCasStatus_t closeMediaCasPvr(void *casObj);
#endif
#ifdef __cplusplus
};
#endif

#endif
