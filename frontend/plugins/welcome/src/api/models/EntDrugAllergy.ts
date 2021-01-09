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
    EntDrugAllergyEdges,
    EntDrugAllergyEdgesFromJSON,
    EntDrugAllergyEdgesFromJSONTyped,
    EntDrugAllergyEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrugAllergy
 */
export interface EntDrugAllergy {
    /**
     * DateTime holds the value of the "dateTime" field.
     * @type {string}
     * @memberof EntDrugAllergy
     */
    dateTime?: string;
    /**
     * 
     * @type {EntDrugAllergyEdges}
     * @memberof EntDrugAllergy
     */
    edges?: EntDrugAllergyEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDrugAllergy
     */
    id?: number;
}

export function EntDrugAllergyFromJSON(json: any): EntDrugAllergy {
    return EntDrugAllergyFromJSONTyped(json, false);
}

export function EntDrugAllergyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrugAllergy {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dateTime': !exists(json, 'dateTime') ? undefined : json['dateTime'],
        'edges': !exists(json, 'edges') ? undefined : EntDrugAllergyEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDrugAllergyToJSON(value?: EntDrugAllergy | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dateTime': value.dateTime,
        'edges': EntDrugAllergyEdgesToJSON(value.edges),
        'id': value.id,
    };
}


