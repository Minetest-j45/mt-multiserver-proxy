package main

import (
	"encoding/json"

	"github.com/anon55555/mt"
)

type ToolGroupCaps struct {
	Uses int16 `json:"uses"`
	MaxLvl int16 `json:"maxlevel"`

	Times map[int16]float32 `json:"times"`
}

type ToolCaps struct {
	NonNil bool `json:"-"`

	AttackCooldown float32 `json:"full_punch_interval`
	MaxDropLvl int16 `json:"max_drop_level"`

	GroupCaps map[string]ToolGroupCaps `json:"groupcaps"`

	DmgGroups map[string]int16 `json:"damage_groups"`

	AttackUses uint16 `json:"punch_attack_uses"`
}

func (t ToolCaps) toMT() mt.ToolCaps {
	tc := mt.ToolCaps{
		NonNil: t.NonNil,
		AttackCooldown: t.AttackCooldown,
		MaxDropLvl: t.MaxDropLvl,
		AttackUses: t.AttackUses,
	}

	for k, v := range t.GroupCaps {
		gc := mt.ToolGroupCaps{
			Name: k,
			Uses: v.Uses,
			MaxLvl: v.MaxLvl,
		}

		for k2, v2 := range v.Times {
			gc.Times = append(gc.Times, mt.DigTime{
				Rating: k2,
				Time: v2,
			})
		}

		tc.GroupCaps = append(tc.GroupCaps, gc)
	}

	for k, v := range t.DmgGroups {
		tc.DmgGroups = append(tc.DmgGroups, mt.Group{
			Name: k,
			Rating: v,
		})
	}

	return tc
}

func (t *ToolCaps) fromMT(tc mt.ToolCaps) {
	t.NonNil = tc.NonNil
	t.AttackCooldown = tc.AttackCooldown
	t.MaxDropLvl = tc.MaxDropLvl
	t.GroupCaps = make(map[string]ToolGroupCaps)
	t.DmgGroups = make(map[string]int16)
	t.AttackUses = tc.AttackUses

	for _, gc := range tc.GroupCaps {
		g := ToolGroupCaps{
			Uses: gc.Uses,
			MaxLvl: gc.MaxLvl,
		}

		for _, dt := range gc.Times {
			g.Times[dt.Rating] = dt.Time
		}

		t.GroupCaps[gc.Name] = g
	}

	for _, g := range tc.DmgGroups {
		t.DmgGroups[g.Name] = g.Rating
	}
}

func (t ToolCaps) SerializeJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(t)
}

func (t *ToolCaps) DeserializeJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(t)
}
