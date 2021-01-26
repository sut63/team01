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
     * Ordercompany holds the value of the ordercompany edge.
     * @type {Array<EntOrder>}
     * @memberof EntCompanyEdges
     */
    ordercompany?: Array<EntOrder>;
}

export function EntCompanyEdgesFromJSON(json: any): EntCompanyEdges {
    return EntCompanyEdgesFromJSONTyped(json, false);
}

export function EntCompanyEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCompanyEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'ordercompany': !exists(json, 'ordercompany') ? undefined : ((json['ordercompany'] as Array<any>).map(EntOrderFromJSON)),
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
        
        'ordercompany': value.ordercompany === undefined ? undefined : ((value.ordercompany as Array<any>).map(EntOrderToJSON)),
    };
}


