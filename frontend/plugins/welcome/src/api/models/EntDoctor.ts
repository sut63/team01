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
    EntDoctorEdges,
    EntDoctorEdgesFromJSON,
    EntDoctorEdgesFromJSONTyped,
    EntDoctorEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDoctor
 */
export interface EntDoctor {
    /**
     * 
     * @type {EntDoctorEdges}
     * @memberof EntDoctor
     */
    edges?: EntDoctorEdges;
    /**
     * Email holds the value of the "email" field.
     * @type {string}
     * @memberof EntDoctor
     */
    email?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDoctor
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntDoctor
     */
    name?: string;
    /**
     * Password holds the value of the "password" field.
     * @type {string}
     * @memberof EntDoctor
     */
    password?: string;
}

export function EntDoctorFromJSON(json: any): EntDoctor {
    return EntDoctorFromJSONTyped(json, false);
}

export function EntDoctorFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDoctor {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntDoctorEdgesFromJSON(json['edges']),
        'email': !exists(json, 'email') ? undefined : json['email'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'password': !exists(json, 'password') ? undefined : json['password'],
    };
}

export function EntDoctorToJSON(value?: EntDoctor | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntDoctorEdgesToJSON(value.edges),
        'email': value.email,
        'id': value.id,
        'name': value.name,
        'password': value.password,
    };
}


