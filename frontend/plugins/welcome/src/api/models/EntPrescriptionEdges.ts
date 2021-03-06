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
    EntDispenseMedicine,
    EntDispenseMedicineFromJSON,
    EntDispenseMedicineFromJSONTyped,
    EntDispenseMedicineToJSON,
    EntDoctor,
    EntDoctorFromJSON,
    EntDoctorFromJSONTyped,
    EntDoctorToJSON,
    EntMedicine,
    EntMedicineFromJSON,
    EntMedicineFromJSONTyped,
    EntMedicineToJSON,
    EntPatientInfo,
    EntPatientInfoFromJSON,
    EntPatientInfoFromJSONTyped,
    EntPatientInfoToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPrescriptionEdges
 */
export interface EntPrescriptionEdges {
    /**
     * 
     * @type {EntDispenseMedicine}
     * @memberof EntPrescriptionEdges
     */
    dispensemedicine?: EntDispenseMedicine;
    /**
     * 
     * @type {EntDoctor}
     * @memberof EntPrescriptionEdges
     */
    prescriptiondoctor?: EntDoctor;
    /**
     * 
     * @type {EntMedicine}
     * @memberof EntPrescriptionEdges
     */
    prescriptionmedicine?: EntMedicine;
    /**
     * 
     * @type {EntPatientInfo}
     * @memberof EntPrescriptionEdges
     */
    prescriptionpatient?: EntPatientInfo;
}

export function EntPrescriptionEdgesFromJSON(json: any): EntPrescriptionEdges {
    return EntPrescriptionEdgesFromJSONTyped(json, false);
}

export function EntPrescriptionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPrescriptionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dispensemedicine': !exists(json, 'Dispensemedicine') ? undefined : EntDispenseMedicineFromJSON(json['Dispensemedicine']),
        'prescriptiondoctor': !exists(json, 'Prescriptiondoctor') ? undefined : EntDoctorFromJSON(json['Prescriptiondoctor']),
        'prescriptionmedicine': !exists(json, 'Prescriptionmedicine') ? undefined : EntMedicineFromJSON(json['Prescriptionmedicine']),
        'prescriptionpatient': !exists(json, 'Prescriptionpatient') ? undefined : EntPatientInfoFromJSON(json['Prescriptionpatient']),
    };
}

export function EntPrescriptionEdgesToJSON(value?: EntPrescriptionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dispensemedicine': EntDispenseMedicineToJSON(value.dispensemedicine),
        'prescriptiondoctor': EntDoctorToJSON(value.prescriptiondoctor),
        'prescriptionmedicine': EntMedicineToJSON(value.prescriptionmedicine),
        'prescriptionpatient': EntPatientInfoToJSON(value.prescriptionpatient),
    };
}


