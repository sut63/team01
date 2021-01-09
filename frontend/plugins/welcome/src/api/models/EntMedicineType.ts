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
    EntMedicineTypeEdges,
    EntMedicineTypeEdgesFromJSON,
    EntMedicineTypeEdgesFromJSONTyped,
    EntMedicineTypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMedicineType
 */
export interface EntMedicineType {
    /**
     * 
     * @type {EntMedicineTypeEdges}
     * @memberof EntMedicineType
     */
    edges?: EntMedicineTypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntMedicineType
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntMedicineType
     */
    name?: string;
}

export function EntMedicineTypeFromJSON(json: any): EntMedicineType {
    return EntMedicineTypeFromJSONTyped(json, false);
}

export function EntMedicineTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMedicineType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntMedicineTypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function EntMedicineTypeToJSON(value?: EntMedicineType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntMedicineTypeEdgesToJSON(value.edges),
        'id': value.id,
        'name': value.name,
    };
}

