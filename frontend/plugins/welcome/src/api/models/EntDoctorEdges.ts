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
    EntPrescription,
    EntPrescriptionFromJSON,
    EntPrescriptionFromJSONTyped,
    EntPrescriptionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDoctorEdges
 */
export interface EntDoctorEdges {
    /**
     * Doctorprescription holds the value of the doctorprescription edge.
     * @type {Array<EntPrescription>}
     * @memberof EntDoctorEdges
     */
    doctorprescription?: Array<EntPrescription>;
}

export function EntDoctorEdgesFromJSON(json: any): EntDoctorEdges {
    return EntDoctorEdgesFromJSONTyped(json, false);
}

export function EntDoctorEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDoctorEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'doctorprescription': !exists(json, 'Doctorprescription') ? undefined : ((json['Doctorprescription'] as Array<any>).map(EntPrescriptionFromJSON)),
    };
}

export function EntDoctorEdgesToJSON(value?: EntDoctorEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'doctorprescription': value.doctorprescription === undefined ? undefined : ((value.doctorprescription as Array<any>).map(EntPrescriptionToJSON)),
    };
}


