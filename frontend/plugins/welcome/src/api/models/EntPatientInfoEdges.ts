/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDrugAllergy,
    EntDrugAllergyFromJSON,
    EntDrugAllergyFromJSONTyped,
    EntDrugAllergyToJSON,
    EntPrescription,
    EntPrescriptionFromJSON,
    EntPrescriptionFromJSONTyped,
    EntPrescriptionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientInfoEdges
 */
export interface EntPatientInfoEdges {
    /**
     * Drugallergys holds the value of the drugallergys edge.
     * @type {Array<EntDrugAllergy>}
     * @memberof EntPatientInfoEdges
     */
    drugallergys?: Array<EntDrugAllergy>;
    /**
     * Patientprescription holds the value of the patientprescription edge.
     * @type {Array<EntPrescription>}
     * @memberof EntPatientInfoEdges
     */
    patientprescription?: Array<EntPrescription>;
}

export function EntPatientInfoEdgesFromJSON(json: any): EntPatientInfoEdges {
    return EntPatientInfoEdgesFromJSONTyped(json, false);
}

export function EntPatientInfoEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientInfoEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'drugallergys': !exists(json, 'drugallergys') ? undefined : ((json['drugallergys'] as Array<any>).map(EntDrugAllergyFromJSON)),
        'patientprescription': !exists(json, 'patientprescription') ? undefined : ((json['patientprescription'] as Array<any>).map(EntPrescriptionFromJSON)),
    };
}

export function EntPatientInfoEdgesToJSON(value?: EntPatientInfoEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'drugallergys': value.drugallergys === undefined ? undefined : ((value.drugallergys as Array<any>).map(EntDrugAllergyToJSON)),
        'patientprescription': value.patientprescription === undefined ? undefined : ((value.patientprescription as Array<any>).map(EntPrescriptionToJSON)),
    };
}


