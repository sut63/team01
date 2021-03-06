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
    EntPatientInfoEdges,
    EntPatientInfoEdgesFromJSON,
    EntPatientInfoEdgesFromJSONTyped,
    EntPatientInfoEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientInfo
 */
export interface EntPatientInfo {
    /**
     * Age holds the value of the "age" field.
     * @type {number}
     * @memberof EntPatientInfo
     */
    age?: number;
    /**
     * CardNumber holds the value of the "cardNumber" field.
     * @type {string}
     * @memberof EntPatientInfo
     */
    cardNumber?: string;
    /**
     * 
     * @type {EntPatientInfoEdges}
     * @memberof EntPatientInfo
     */
    edges?: EntPatientInfoEdges;
    /**
     * Gender holds the value of the "gender" field.
     * @type {string}
     * @memberof EntPatientInfo
     */
    gender?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatientInfo
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntPatientInfo
     */
    name?: string;
}

export function EntPatientInfoFromJSON(json: any): EntPatientInfo {
    return EntPatientInfoFromJSONTyped(json, false);
}

export function EntPatientInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'age': !exists(json, 'age') ? undefined : json['age'],
        'cardNumber': !exists(json, 'cardNumber') ? undefined : json['cardNumber'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientInfoEdgesFromJSON(json['edges']),
        'gender': !exists(json, 'gender') ? undefined : json['gender'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function EntPatientInfoToJSON(value?: EntPatientInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'age': value.age,
        'cardNumber': value.cardNumber,
        'edges': EntPatientInfoEdgesToJSON(value.edges),
        'gender': value.gender,
        'id': value.id,
        'name': value.name,
    };
}


