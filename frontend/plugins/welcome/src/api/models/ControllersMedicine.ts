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
/**
 * 
 * @export
 * @interface ControllersMedicine
 */
export interface ControllersMedicine {
    /**
     * 
     * @type {number}
     * @memberof ControllersMedicine
     */
    levelOfDangerousID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersMedicine
     */
    medicineTypeID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersMedicine
     */
    unitOfMedicineID?: number;
}

export function ControllersMedicineFromJSON(json: any): ControllersMedicine {
    return ControllersMedicineFromJSONTyped(json, false);
}

export function ControllersMedicineFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersMedicine {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'levelOfDangerousID': !exists(json, 'levelOfDangerousID') ? undefined : json['levelOfDangerousID'],
        'medicineTypeID': !exists(json, 'medicineTypeID') ? undefined : json['medicineTypeID'],
        'unitOfMedicineID': !exists(json, 'unitOfMedicineID') ? undefined : json['unitOfMedicineID'],
    };
}

export function ControllersMedicineToJSON(value?: ControllersMedicine | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'levelOfDangerousID': value.levelOfDangerousID,
        'medicineTypeID': value.medicineTypeID,
        'unitOfMedicineID': value.unitOfMedicineID,
    };
}


