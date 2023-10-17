package com.droidlogic.jcas;

import static com.droidlogic.jcas.CasManager.EMM_FILTERING_ON_TUNE;
import static com.droidlogic.jcas.CasManager.START_CAS_SERVICE_ON_BOOT;

import com.droidlogic.jcas.vendor.widevine.WidevineCasProjectListener;
import com.droidlogic.jcas.vendor.widevine.WidevineCasWrapperFactory;

/**
 * A factory class containing method for creating a JCas for the current platform
 */
public class JCasBuilder {
    /**
     * Project can choose each Cas vendor's CasWrapperFactory in project library.
     * ex) If WidevineCas case, need to change
     * 'WidevineCasWrapperFactory' to '{vendor}CasWrapperFactory'
     */
    public static CasWrapperFactory[] createFactory(String config) {
        return new CasWrapperFactory[]{new WidevineCasWrapperFactory(config)};
    }

    /**
     * Project can choose a each Cas vendor's CasProjectListener in project library.
     * ex) If WidevineCas case, need to change
     * 'WidevineCasProjectListener' to '{vendor}CasProjectListener'
     */
    public static CasProjectListener createListener() {
        return new WidevineCasProjectListener();
    }

    /**
     * Project can choose a Emm filtering mode.
     * EMM_FILTERING_ON_TUNE:
     *  - EMM filtering lifetime is same as each Channel
     *  - It used only live connection to send CAT/EMM data.
     * EMM_FILTERING_ALWAYS:
     *  - EMM filtering lifetime is same as JCas Service
     *  - It used only default connection using defaultCasWrapper to send CAT/EMM data.
     *
     * Default Mode: EMM_FILTERING_ON_TUNE
     */
    public static int getEmmFilteringMode() {
        return EMM_FILTERING_ON_TUNE;
    }

    /**
     * Project can choose the content rating source
     * false: default, use only SI's rating
     * true: use Cas's content rating event on every scramble channels
     */
    public static boolean getCasContentRatingEnabled() {
        return false;
    }

    /**
     * The Project can choose when to start CasService.
     *
     * START_CAS_SERVICE_ON_BOOT:
     * The casService started when received the BOOT_COMPLETED intent.
     *
     * START_CAS_SERVICE_ON_REQUEST:
     * The casService started when receiving the below intent from the TV app.
     *      - Intent name: com.droidlogic.jcas.ACTION_START_CAS_SERVICE
     *      - Permission: com.droidlogic.jcas.permission.START_CAS_SERVICE
     *
     * Example) Intent startCasService = new Intent(INTENT_START_CAS_SERVICE);
     *          startCasService.setPackage("com.google.android.tv.dtvinput");
     *          Context.sendBroadcast(startCasService,
     *          "com.droidlogic.jcas.permission.START_CAS_SERVICE");
     *
     * Default mode: START_CAS_SERVICE_ON_BOOT
     */
    public static int getStartCasServiceMode() {
        return START_CAS_SERVICE_ON_BOOT;
    }
}
