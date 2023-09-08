/*
 * Copyright (c) 2023 Amlogic, Inc. All rights reserved.
 *
 * This source code is subject to the terms and conditions defined in the
 * file 'LICENSE' which is part of this source code package.
 *
 * Description:
 */

#ifndef _WVCAS_TYPES_H_
#define _WVCAS_TYPES_H_

#define WV_CA_ID 0x4AD4
#define MAX_ECM_SECTION_LEN 1024

typedef enum
{
    WVCAS_SUCCESS,
    WVCAS_ERROR,
    WVCAS_ERROR_INIT_FAILED,
    WVCAS_ERROR_OPEN_SESSION_FAILURE,
    WVCAS_ERROR_NOT_IMPLEMENTED,
} WvCasResult;

typedef struct __tagCAS_STREAM_INFO {
    unsigned int        ca_system_id;
    uint16_t            desc_num;
    unsigned int        ecm_pid[4];
    uint16_t            audio_pid;
    uint16_t            video_pid;
    int                 audio_channel;
    int                 video_channel;
    bool                av_diff_ecm;
    uint8_t             *private_data;
    unsigned int        pri_data_len;
    void                *headers;
} CasStreamInfo;

typedef struct __tagCAS_INIT_HEADERS {
	char *license_url;
	char *provision_url;
	char *request_header;
	char *request_body;
	char *content_type;
} CasInitHeaders;

typedef struct __tagECM_INFO {
	int first_set_ecm;
	int ecm_data_size; //first ts packet header + ecm section header + ecm section
	int ecm_copy_size;
	bool found_ecm_payload;
	uint8_t pre_ecm_data[MAX_ECM_SECTION_LEN];
	uint8_t ecm_data[MAX_ECM_SECTION_LEN];
} EcmInfo;

#endif
