// Code generated by tygo. DO NOT EDIT.
/* eslint-disable @typescript-eslint/no-unnecessary-type-constraint */
/* eslint-disable spaced-comment */
import type { ApiDataResponse } from './common'


//////////
// source: slot_viz.go

/**
 * ------------------------------------------------------------
 * Slot Viz
 */
export interface VDBSlotVizPassiveDuty {
    pending_count: number /* uint64 */;
    success_count: number /* uint64 */;
    failed_count: number /* uint64 */;
}
export interface VDBSlotVizActiveDuty {
    status: 'success' | 'failed' | 'scheduled';
    validator: number /* uint64 */;
    /**
     * 		If the duty is a proposal & it's successful, the duty_object is the proposed block
     * 		If the duty is a proposal & it failed/scheduled, the duty_object is the slot
     * 		If the duty is a slashing & it's successful, the duty_object is the validator you slashed
     * 		If the duty is a slashing & it failed, the duty_object is your validator that was slashed
     */
    duty_object: number /* uint64 */;
}
export interface VDBSlotVizSlot {
    slot: number /* uint64 */;
    status: 'proposed' | 'missed' | 'scheduled' | 'orphaned';
    attestations?: VDBSlotVizPassiveDuty;
    sync?: VDBSlotVizPassiveDuty;
    proposals?: VDBSlotVizActiveDuty[];
    slashing?: VDBSlotVizActiveDuty[];
}
export interface SlotVizEpoch {
    epoch: number /* uint64 */;
    state?: 'head' | 'finalized' | 'scheduled'; // only on landing page
    progress?: number /* float64 */; // only on landing page
    slots?: VDBSlotVizSlot[]; // only on dashboard page
}
export type InternalGetValidatorDashboardSlotVizResponse = ApiDataResponse<SlotVizEpoch[]>;
