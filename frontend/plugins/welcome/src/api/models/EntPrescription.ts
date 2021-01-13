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
    EntPrescriptionEdges,
    EntPrescriptionEdgesFromJSON,
    EntPrescriptionEdgesFromJSONTyped,
    EntPrescriptionEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPrescription
 */
export interface EntPrescription {
    /**
     * Value holds the value of the "Value" field.
     * @type {number}
     * @memberof EntPrescription
     */
    value?: number;
    /**
     * 
     * @type {EntPrescriptionEdges}
     * @memberof EntPrescription
     */
    edges?: EntPrescriptionEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPrescription
     */
    id?: number;
}

export function EntPrescriptionFromJSON(json: any): EntPrescription {
    return EntPrescriptionFromJSONTyped(json, false);
}

export function EntPrescriptionFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPrescription {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'value': !exists(json, 'Value') ? undefined : json['Value'],
        'edges': !exists(json, 'edges') ? undefined : EntPrescriptionEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPrescriptionToJSON(value?: EntPrescription | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Value': value.value,
        'edges': EntPrescriptionEdgesToJSON(value.edges),
        'id': value.id,
    };
}


