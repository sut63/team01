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
    EntBillEdges,
    EntBillEdgesFromJSON,
    EntBillEdgesFromJSONTyped,
    EntBillEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBill
 */
export interface EntBill {
    /**
     * Amount holds the value of the "amount" field.
     * @type {number}
     * @memberof EntBill
     */
    amount?: number;
    /**
     * Annotation holds the value of the "annotation" field.
     * @type {string}
     * @memberof EntBill
     */
    annotation?: string;
    /**
     * 
     * @type {EntBillEdges}
     * @memberof EntBill
     */
    edges?: EntBillEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBill
     */
    id?: number;
    /**
     * Payer holds the value of the "payer" field.
     * @type {string}
     * @memberof EntBill
     */
    payer?: string;
}

export function EntBillFromJSON(json: any): EntBill {
    return EntBillFromJSONTyped(json, false);
}

export function EntBillFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBill {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'annotation': !exists(json, 'annotation') ? undefined : json['annotation'],
        'edges': !exists(json, 'edges') ? undefined : EntBillEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'payer': !exists(json, 'payer') ? undefined : json['payer'],
    };
}

export function EntBillToJSON(value?: EntBill | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'amount': value.amount,
        'annotation': value.annotation,
        'edges': EntBillEdgesToJSON(value.edges),
        'id': value.id,
        'payer': value.payer,
    };
}


