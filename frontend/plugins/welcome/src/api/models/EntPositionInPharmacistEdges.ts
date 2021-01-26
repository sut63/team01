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
    EntPharmacist,
    EntPharmacistFromJSON,
    EntPharmacistFromJSONTyped,
    EntPharmacistToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPositionInPharmacistEdges
 */
export interface EntPositionInPharmacistEdges {
    /**
     * Pharmacist holds the value of the pharmacist edge.
     * @type {Array<EntPharmacist>}
     * @memberof EntPositionInPharmacistEdges
     */
    pharmacist?: Array<EntPharmacist>;
}

export function EntPositionInPharmacistEdgesFromJSON(json: any): EntPositionInPharmacistEdges {
    return EntPositionInPharmacistEdgesFromJSONTyped(json, false);
}

export function EntPositionInPharmacistEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPositionInPharmacistEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pharmacist': !exists(json, 'Pharmacist') ? undefined : ((json['Pharmacist'] as Array<any>).map(EntPharmacistFromJSON)),
    };
}

export function EntPositionInPharmacistEdgesToJSON(value?: EntPositionInPharmacistEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pharmacist': value.pharmacist === undefined ? undefined : ((value.pharmacist as Array<any>).map(EntPharmacistToJSON)),
    };
}


