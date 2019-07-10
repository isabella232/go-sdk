/****************************************************************************
 * Copyright 2019, Optimizely, Inc. and contributors                        *
 *                                                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");          *
 * you may not use this file except in compliance with the License.         *
 * You may obtain a copy of the License at                                  *
 *                                                                          *
 *    http://www.apache.org/licenses/LICENSE-2.0                            *
 *                                                                          *
 * Unless required by applicable law or agreed to in writing, software      *
 * distributed under the License is distributed on an "AS IS" BASIS,        *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
 * See the License for the specific language governing permissions and      *
 * limitations under the License.                                           *
 ***************************************************************************/

package decision

import (
	"github.com/optimizely/go-sdk/optimizely/decision/evaluator"
	"github.com/optimizely/go-sdk/optimizely/entities"
)

// ExperimentTargetingService makes a decision using audience targeting
type ExperimentTargetingService struct {
	audienceEvaluator evaluator.AudienceEvaluator
}

// NewExperimentTargetingService returns a new instance of ExperimentTargetingService
func NewExperimentTargetingService() *ExperimentTargetingService {
	return &ExperimentTargetingService{
		audienceEvaluator: evaluator.NewTypedAudienceEvaluator(),
	}
}

// GetDecision applies audience targeting from the given experiment to the given user. The only decision it makes is whether to exclude the user from the experiment.
func (s ExperimentTargetingService) GetDecision(decisionContext ExperimentDecisionContext, userContext entities.UserContext) (ExperimentDecision, error) {
	experimentDecision := ExperimentDecision{}
	experiment := decisionContext.Experiment
	if len(experiment.AudienceIds) > 0 {
		experimentAudience := decisionContext.AudienceMap[experiment.AudienceIds[0]]
		evalResult := s.audienceEvaluator.Evaluate(experimentAudience, userContext)
		if evalResult == false {
			// user not targeted for experiment, return an empty variation
			experimentDecision.DecisionMade = true
			experimentDecision.Variation = entities.Variation{}
			return experimentDecision, nil
		}
	}
	// user passes audience targeting, can move on to the next decision maker
	return experimentDecision, nil
}