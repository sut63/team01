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
    EntOrder,
    EntOrderFromJSON,
    EntOrderFromJSONTyped,
    EntOrderToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCompanyEdges
 */
export interface EntCompanyEdges {
    /**
     * Order holds the value of the order edge.
     * @type {Array<EntOrder>}
     * @memberof EntCompanyEdges
     */
    order?: Array<EntOrder>;
}

export function EntCompanyEdgesFromJSON(json: any): EntCompanyEdges {
    return EntCompanyEdgesFromJSONTyped(json, false);
}

export function EntCompanyEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCompanyEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'order': !exists(json, 'order') ? undefined : ((json['order'] as Array<any>).map(EntOrderFromJSON)),
    };
}

export function EntCompanyEdgesToJSON(value?: EntCompanyEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'order': value.order === undefined ? undefined : ((value.order as Array<any>).map(EntOrderToJSON)),
    };
}

