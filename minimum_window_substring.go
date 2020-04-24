// Copyright (c) 2012-2018 Grabtaxi Holdings PTE LTD (GRAB), All Rights Reserved. NOTICE: All information contained herein
// is, and remains the property of GRAB. The intellectual and technical concepts contained herein are confidential, proprietary
// and controlled by GRAB and may be covered by patents, patents in process, and are protected by trade secret or copyright law.
//
// You are strictly forbidden to copy, download, store (in any medium), transmit, disseminate, adapt or change this material
// in any way unless prior written permission is obtained from GRAB. Access to the source code contained herein is hereby
// forbidden to anyone except current GRAB employees or contractors with binding Confidentiality and Non-disclosure agreements
// explicitly covering such access.
//
// The copyright notice above does not evidence any actual or intended publication or disclosure of this source code,
// which includes information that is confidential and/or proprietary, and is a trade secret, of GRAB.
//
// ANY REPRODUCTION, MODIFICATION, DISTRIBUTION, PUBLIC PERFORMANCE, OR PUBLIC DISPLAY OF OR THROUGH USE OF THIS SOURCE
// CODE WITHOUT THE EXPRESS WRITTEN CONSENT OF GRAB IS STRICTLY PROHIBITED, AND IN VIOLATION OF APPLICABLE LAWS AND
// INTERNATIONAL TREATIES. THE RECEIPT OR POSSESSION OF THIS SOURCE CODE AND/OR RELATED INFORMATION DOES NOT CONVEY
// OR IMPLY ANY RIGHTS TO REPRODUCE, DISCLOSE OR DISTRIBUTE ITS CONTENTS, OR TO MANUFACTURE, USE, OR SELL ANYTHING
// THAT IT MAY DESCRIBE, IN WHOLE OR IN PART.

package algos

import "container/list"

/*
Question: https://leetcode.com/problems/minimum-window-substring/solution/
 */

// minWindow assumes characters in t are unique
func minWindow(s string, t string) string {
	// build map of characteres in t
	tmap := map[int32]bool{}
	num_elements := 0
	for _, ch := range t{
		if _,ok := tmap[ch]; !ok{
			num_elements ++
			tmap[ch] = true
		}
	}

	// loop through s
	last_pos_map := map[int32]*list.Element{} // maps from character to the corresponding element in last_pos_list
	last_pos_list := list.New() // records the last-seen position (ordered asc) of each character
	elements_found := 0
	minimum_length := int(1e6)
	minimum_window := ""

	for i, ch := range s{
		// we are only interested in the characters in t
		if _,ok := tmap[ch]; !ok{
			continue
		}

		// we couldn't start calculating minimum window length until we encounter all the characters in t at least once
		if elements_found < num_elements {
			if _,ok := last_pos_map[ch]; !ok{
				elements_found ++;

				// this is the first valid window
				if elements_found == num_elements{
					earliestElemPos := 0
					if last_pos_list.Front() != nil{
						earliestElemPos = last_pos_list.Front().Value.(int)
					}
					minimum_length = i - earliestElemPos
					minimum_window = s[earliestElemPos:i+1]
				}
			} else{
				last_pos_list.Remove(last_pos_map[ch])
			}
			last_pos_map[ch] = last_pos_list.PushBack(i);

		// we have already found one valid window, next we will update minimum window if we encounter new smaller windows
		} else{
			earliestElem := last_pos_list.Front().Value.(int)
			// when we encounter the earliest element again, it means a new window is found
			// e.g. a ... b .... c ... a, when 'a' is seen again, the b,c and the later a form a new window
			if string(s[earliestElem]) == string(ch) {
				last_pos_list.Remove(last_pos_list.Front())
				secondEarliestPos := last_pos_list.Front().Value.(int) // position of b
				if i - secondEarliestPos < minimum_length{
					minimum_length = i - secondEarliestPos
					minimum_window = s[secondEarliestPos:i+1]
				}
				last_pos_map[ch] = last_pos_list.PushBack(i);
			} else {
				last_pos_list.Remove(last_pos_map[ch])
				last_pos_map[ch] = last_pos_list.PushBack(i);
			}
		}
	}

	return minimum_window
}
