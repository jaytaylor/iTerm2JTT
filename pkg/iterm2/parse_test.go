package iterm2

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		content  string // Source file content.
		expected Tips   // Number of expected tips to be parsed from content.
	}{
		{
			content:  ``,
			expected: Tips{},
		},
		{
			content:  `            @"001": @{ kTipTitleKey: @"Tip of the Day,`,
			expected: Tips{Tip{ID: 1}},
		},
		{
			content:  `            @"002": @{ kTipTitleKey: "Tip of the Day",`,
			expected: Tips{Tip{ID: 2}},
		},
		{
			content:  `            "003": @{ kTipTitleKey: @"Tip of the Day",`,
			expected: Tips{Tip{Title: "Tip of the Day"}},
		},
		{
			content:  `            @0004: @{ kTipTitleKey: @"Tip of the Day",`,
			expected: Tips{Tip{Title: "Tip of the Day"}},
		},
		{
			content:  `            00005: @{ kTipTitleKey: @"Tip of the Day",`,
			expected: Tips{Tip{Title: "Tip of the Day"}},
		},
		{
			content:  `            @"0006": @{ kTipTitleKey: @"Tip of the Day",`,
			expected: Tips{Tip{ID: 6, Title: "Tip of the Day"}},
		},
		{
			content: `+ (NSDictionary *)allTips {
  // The keys in this dictionary are saved in user defaults and should not be changed or
  // recycled, or users will see the same tip more than once.
  return @{
    // Big new features
            @"010": @{ kTipTitleKey: @"Tip of the Day",
                        kTipBodyKey: @"This window shows the iTerm2 tip of the day. It’ll appear every 24 hours to let you know about new features and hidden secrets. Hit “More Options” to view more tips or to stop getting them altogether." },
            @"0100": @{ kTipTitleKey: @"Shell Integration",
                         kTipBodyKey: @"The big new feature of iTerm2 version 3 is Shell Integration. Click “Learn More” for all the details.",
                          kTipUrlKey: @"https://iterm2.com/shell_integration.html" },`,
			expected: Tips{
				Tip{
					ID:    10,
					Title: "Tip of the Day",
					Body:  "This window shows the iTerm2 tip of the day. It’ll appear every 24 hours to let you know about new features and hidden secrets. Hit “More Options” to view more tips or to stop getting them altogether.",
				},
				Tip{
					ID:    100,
					Title: "Shell Integration",
					Body:  "The big new feature of iTerm2 version 3 is Shell Integration. Click “Learn More” for all the details.",
					URL:   "https://iterm2.com/shell_integration.html",
				},
			},
		},
		{
			content: `//
//  iTermTipData.m
//  iTerm2
//
//  Created by George Nachman on 6/19/15.
//
//

#import "iTermTipData.h"
#import "iTermTip.h"

@implementation iTermTipData

+ (NSDictionary *)allTips {
  // The keys in this dictionary are saved in user defaults and should not be changed or
  // recycled, or users will see the same tip more than once.
  return @{
    // Big new features
}

@end`,
			expected: Tips{},
		},
		{
			content: `//
//  iTermTipData.m
//  iTerm2
//
//  Created by George Nachman on 6/19/15.
//
//

#import "iTermTipData.h"
#import "iTermTip.h"

@implementation iTermTipData

+ (NSDictionary *)allTips {
  // The keys in this dictionary are saved in user defaults and should not be changed or
  // recycled, or users will see the same tip more than once.
  return @{
    // Big new features
            @"009": @{ kTipTitleKey: @"Tip of the Day",
                       kTipBodyKey: @"This window shows the iTerm2 tip of the day. It’ll appear every 24 hours to let you know about new features and hidden secrets. Hit “More Options” to view more tips or to stop getting them altogether." },
            @"0099": @{ kTipTitleKey: @"Shell Integration",
                        kTipBodyKey: @"The big new feature of iTerm2 version 3 is Shell Integration. Click “Learn More” for all the details.",
                        kTipUrlKey: @"https://iterm2.com/shell_integration.html" },

            @"0999": @{ kTipTitleKey: @"Timestamps",
                        kTipBodyKey: @"“View > Show Timestamps” shows the time (and date, if appropriate) when each line was last modified." },

            @"9971": @{ kTipTitleKey: @"Buried Sessions",
                        kTipBodyKey: @"You can “bury” a session with “Session > Bury Session.” It remains hidden until you restore it by selecting it from “Session > Buried Sessions > Your session.”" },

            };
}

@end`,
			expected: Tips{
				Tip{
					ID:    9,
					Title: "Tip of the Day",
					Body:  "This window shows the iTerm2 tip of the day. It’ll appear every 24 hours to let you know about new features and hidden secrets. Hit “More Options” to view more tips or to stop getting them altogether.",
				},
				Tip{
					ID:    99,
					Title: "Shell Integration",
					Body:  "The big new feature of iTerm2 version 3 is Shell Integration. Click “Learn More” for all the details.",
					URL:   "https://iterm2.com/shell_integration.html",
				},
				Tip{
					ID:    999,
					Title: "Timestamps",
					Body:  "“View > Show Timestamps” shows the time (and date, if appropriate) when each line was last modified.",
				},
				Tip{
					ID:    9971,
					Title: "Buried Sessions",
					Body:  "You can “bury” a session with “Session > Bury Session.” It remains hidden until you restore it by selecting it from “Session > Buried Sessions > Your session.”",
				},
			},
		},
		{
			content: `+ (NSDictionary *)allTips {
    // Big new features
            @"909": @{ kTipUrlKey: @"https://iterm2.com/shell_integration.html",
                       kTipBodyKey: @"The big new feature of iTerm2 version 3 is Shell Integration. Click “Learn More” for all the details.",
                       kTipTitleKey: @"Shell Integration" },
            };
}

@end`,
			expected: Tips{
				Tip{
					ID:    909,
					Title: "Shell Integration",
					Body:  "The big new feature of iTerm2 version 3 is Shell Integration. Click “Learn More” for all the details.",
					URL:   "https://iterm2.com/shell_integration.html",
				},
			},
		},
	}
	for i, testCase := range testCases {
		r := strings.NewReader(testCase.content)
		tips, err := Parse(r)
		if err != nil {
			t.Errorf("[test-case #%v] Unexpected error parsing test-case: %s\ncontent fragment: >>> %v <<<\ntips: %+v", i, err, strings.Split(testCase.content[0:50], "\n")[0], tips)
		}
		if !reflect.DeepEqual(tips, testCase.expected) {
			t.Errorf("[test-case #%v] Expected tips=%v but actual=%v\ncontent fragment: >>> %v <<<", i, testCase.expected, tips, strings.Split(testCase.content[0:50], "\n")[0])
		}
	}
}
