package com.droidlogic.jcas;

import static com.droidlogic.jcas.CasManager.EMM_FILTERING_ON_TUNE;

import com.droidlogic.jcas.vendor.widevine.WidevineCasProjectListener;
import com.droidlogic.jcas.vendor.widevine.WidevineCasWrapperFactory;

/**
 * A factory class containing method for creating a JCas for the current platform
 */
public class JCasBuilder {
    /**
     * Project can choose each Cas vendor's CasWrapperFactory in project library.
     * ex) If not WidevineCas case, need to change
     * 'WidevineCasWrapperFactory' to '{vendor}CasWrapperFactory'
     */
    public static CasWrapperFactory[] createFactory(String config) {
        return new CasWrapperFactory[]{new WidevineCasWrapperFactory(config)};
    }

    /**
     * Project can choose a each Cas vendor's CasProjectListener in project library.
     * ex) If not WidevineCas case, need to change
     * 'WidevineCasProjectListener' to '{vendor}CasProjectListener'
     */
    public static CasProjectListener createListener() {
        return new WidevineCasProjectListener();
    }

    /**
     * Get the Emm filtering mode.
     * EMM_FILTERING_ON_TUNE:
     *  - EMM filtering lifetime is same as each Channel
     *  - It used only live connection to send CAT/EMM data.
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
}
