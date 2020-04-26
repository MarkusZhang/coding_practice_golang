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

import (
	"math"
)

/*
problem link: https://leetcode.com/problems/super-egg-drop/
 */

// superEggDrop solves the problem with complexity O(eggs * floors^2)
func superEggDrop(eggs int, floors int) int {
	// edge case
	if eggs == 1{
		return floors
	}
	if floors <= 1{
		return 1
	}

	// create DP matrix, with x-axis as number of eggs, y-axis as number of floors
	// m[x][y] represents the minimum number of eggs to drop in the worst case if we have x eggs and y floors
	m := make([][]int, eggs+1);
	for i := range m{
		m[i] = make([]int,floors+1)
	}

	// initialize DP matrix
	for i:=0;i<eggs+1;i++{
		m[i][0] = 0
		// if there is only one floor, we just need to drop once
		m[i][1] = 1
	}
	for j:=0;j<floors+1;j++{
		// if we only have one egg, then # drops required = # floors
		m[1][j] = j
	}

	// fill the matrix using derivation
	for e:=2;e<=eggs;e++{
		for f:=2;f<=floors;f++{
			m[e][f] = calcTransition(m,e,f)
		}
	}
	return m[eggs][floors]
}


// calcTransition calculates value for m[e][f] using the transition function
// m[e][f] = min{ max(m[e-1][k-1], m[e][f-k]) } + 1  where 1 <= k <= f
func calcTransition(m [][]int, e int, f int) int{
	minVal := f
	for k:=1;k<=f;k++{
		val := int(math.Max(
				float64(m[e-1][k-1]),
				float64(m[e][f-k]),
			),
		)
		if val < minVal{
			minVal = val
		}
	}
	return minVal + 1
}
