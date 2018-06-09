// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testEquities(t *testing.T) {
	t.Parallel()

	query := Equities(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testEquitiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = equity.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEquitiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Equities(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEquitiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := EquitySlice{equity}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testEquitiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := EquityExists(tx, equity.InstrumentID)
	if err != nil {
		t.Errorf("Unable to check if Equity exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EquityExistsG to return true, but got false.")
	}
}
func testEquitiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	equityFound, err := FindEquity(tx, equity.InstrumentID)
	if err != nil {
		t.Error(err)
	}

	if equityFound == nil {
		t.Error("want a record, got nil")
	}
}
func testEquitiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Equities(tx).Bind(equity); err != nil {
		t.Error(err)
	}
}

func testEquitiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Equities(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEquitiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equityOne := &Equity{}
	equityTwo := &Equity{}
	if err = randomize.Struct(seed, equityOne, equityDBTypes, false, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}
	if err = randomize.Struct(seed, equityTwo, equityDBTypes, false, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equityOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = equityTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Equities(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEquitiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	equityOne := &Equity{}
	equityTwo := &Equity{}
	if err = randomize.Struct(seed, equityOne, equityDBTypes, false, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}
	if err = randomize.Struct(seed, equityTwo, equityDBTypes, false, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equityOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = equityTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func equityBeforeInsertHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func equityAfterInsertHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func equityAfterSelectHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func equityBeforeUpdateHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func equityAfterUpdateHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func equityBeforeDeleteHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func equityAfterDeleteHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func equityBeforeUpsertHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func equityAfterUpsertHook(e boil.Executor, o *Equity) error {
	*o = Equity{}
	return nil
}

func testEquitiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Equity{}
	o := &Equity{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, equityDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Equity object: %s", err)
	}

	AddEquityHook(boil.BeforeInsertHook, equityBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	equityBeforeInsertHooks = []EquityHook{}

	AddEquityHook(boil.AfterInsertHook, equityAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	equityAfterInsertHooks = []EquityHook{}

	AddEquityHook(boil.AfterSelectHook, equityAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	equityAfterSelectHooks = []EquityHook{}

	AddEquityHook(boil.BeforeUpdateHook, equityBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	equityBeforeUpdateHooks = []EquityHook{}

	AddEquityHook(boil.AfterUpdateHook, equityAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	equityAfterUpdateHooks = []EquityHook{}

	AddEquityHook(boil.BeforeDeleteHook, equityBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	equityBeforeDeleteHooks = []EquityHook{}

	AddEquityHook(boil.AfterDeleteHook, equityAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	equityAfterDeleteHooks = []EquityHook{}

	AddEquityHook(boil.BeforeUpsertHook, equityBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	equityBeforeUpsertHooks = []EquityHook{}

	AddEquityHook(boil.AfterUpsertHook, equityAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	equityAfterUpsertHooks = []EquityHook{}
}
func testEquitiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEquitiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx, equityColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEquityToOneInstrumentUsingInstrument(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Equity
	var foreign Instrument

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, equityDBTypes, false, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, instrumentDBTypes, false, instrumentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Instrument struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.InstrumentID = foreign.InstrumentID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Instrument(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.InstrumentID != foreign.InstrumentID {
		t.Errorf("want: %v, got %v", foreign.InstrumentID, check.InstrumentID)
	}

	slice := EquitySlice{&local}
	if err = local.L.LoadInstrument(tx, false, (*[]*Equity)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Instrument == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Instrument = nil
	if err = local.L.LoadInstrument(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Instrument == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEquityToOneInstrumentUsingPayCurrency(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Equity
	var foreign Instrument

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, instrumentDBTypes, false, instrumentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Instrument struct: %s", err)
	}

	local.PayCurrencyID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PayCurrencyID.Int = foreign.InstrumentID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PayCurrency(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.InstrumentID != foreign.InstrumentID {
		t.Errorf("want: %v, got %v", foreign.InstrumentID, check.InstrumentID)
	}

	slice := EquitySlice{&local}
	if err = local.L.LoadPayCurrency(tx, false, (*[]*Equity)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.PayCurrency == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PayCurrency = nil
	if err = local.L.LoadPayCurrency(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PayCurrency == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEquityToOneInstrumentUsingExerciseCurrency(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Equity
	var foreign Instrument

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, instrumentDBTypes, false, instrumentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Instrument struct: %s", err)
	}

	local.ExerciseCurrencyID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ExerciseCurrencyID.Int = foreign.InstrumentID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ExerciseCurrency(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.InstrumentID != foreign.InstrumentID {
		t.Errorf("want: %v, got %v", foreign.InstrumentID, check.InstrumentID)
	}

	slice := EquitySlice{&local}
	if err = local.L.LoadExerciseCurrency(tx, false, (*[]*Equity)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.ExerciseCurrency == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ExerciseCurrency = nil
	if err = local.L.LoadExerciseCurrency(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ExerciseCurrency == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEquityToOneInstrumentUsingCompanyCurrency(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Equity
	var foreign Instrument

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, instrumentDBTypes, false, instrumentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Instrument struct: %s", err)
	}

	local.CompanyCurrencyID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.CompanyCurrencyID.Int = foreign.InstrumentID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.CompanyCurrency(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.InstrumentID != foreign.InstrumentID {
		t.Errorf("want: %v, got %v", foreign.InstrumentID, check.InstrumentID)
	}

	slice := EquitySlice{&local}
	if err = local.L.LoadCompanyCurrency(tx, false, (*[]*Equity)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.CompanyCurrency == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CompanyCurrency = nil
	if err = local.L.LoadCompanyCurrency(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.CompanyCurrency == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEquityToOneSetOpInstrumentUsingInstrument(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Equity
	var b, c Instrument

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, equityDBTypes, false, strmangle.SetComplement(equityPrimaryKeyColumns, equityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Instrument{&b, &c} {
		err = a.SetInstrument(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Instrument != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Equity != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.InstrumentID != x.InstrumentID {
			t.Error("foreign key was wrong value", a.InstrumentID)
		}

		if exists, err := EquityExists(tx, a.InstrumentID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testEquityToOneSetOpInstrumentUsingPayCurrency(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Equity
	var b, c Instrument

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, equityDBTypes, false, strmangle.SetComplement(equityPrimaryKeyColumns, equityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Instrument{&b, &c} {
		err = a.SetPayCurrency(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PayCurrency != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PayCurrencyEquities[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PayCurrencyID.Int != x.InstrumentID {
			t.Error("foreign key was wrong value", a.PayCurrencyID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PayCurrencyID.Int))
		reflect.Indirect(reflect.ValueOf(&a.PayCurrencyID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PayCurrencyID.Int != x.InstrumentID {
			t.Error("foreign key was wrong value", a.PayCurrencyID.Int, x.InstrumentID)
		}
	}
}

func testEquityToOneRemoveOpInstrumentUsingPayCurrency(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Equity
	var b Instrument

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, equityDBTypes, false, strmangle.SetComplement(equityPrimaryKeyColumns, equityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPayCurrency(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePayCurrency(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.PayCurrency(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.PayCurrency != nil {
		t.Error("R struct entry should be nil")
	}

	if a.PayCurrencyID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PayCurrencyEquities) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testEquityToOneSetOpInstrumentUsingExerciseCurrency(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Equity
	var b, c Instrument

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, equityDBTypes, false, strmangle.SetComplement(equityPrimaryKeyColumns, equityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Instrument{&b, &c} {
		err = a.SetExerciseCurrency(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ExerciseCurrency != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ExerciseCurrencyEquities[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ExerciseCurrencyID.Int != x.InstrumentID {
			t.Error("foreign key was wrong value", a.ExerciseCurrencyID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ExerciseCurrencyID.Int))
		reflect.Indirect(reflect.ValueOf(&a.ExerciseCurrencyID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ExerciseCurrencyID.Int != x.InstrumentID {
			t.Error("foreign key was wrong value", a.ExerciseCurrencyID.Int, x.InstrumentID)
		}
	}
}

func testEquityToOneRemoveOpInstrumentUsingExerciseCurrency(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Equity
	var b Instrument

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, equityDBTypes, false, strmangle.SetComplement(equityPrimaryKeyColumns, equityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetExerciseCurrency(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveExerciseCurrency(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ExerciseCurrency(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ExerciseCurrency != nil {
		t.Error("R struct entry should be nil")
	}

	if a.ExerciseCurrencyID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ExerciseCurrencyEquities) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testEquityToOneSetOpInstrumentUsingCompanyCurrency(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Equity
	var b, c Instrument

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, equityDBTypes, false, strmangle.SetComplement(equityPrimaryKeyColumns, equityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Instrument{&b, &c} {
		err = a.SetCompanyCurrency(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CompanyCurrency != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CompanyCurrencyEquities[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CompanyCurrencyID.Int != x.InstrumentID {
			t.Error("foreign key was wrong value", a.CompanyCurrencyID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CompanyCurrencyID.Int))
		reflect.Indirect(reflect.ValueOf(&a.CompanyCurrencyID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CompanyCurrencyID.Int != x.InstrumentID {
			t.Error("foreign key was wrong value", a.CompanyCurrencyID.Int, x.InstrumentID)
		}
	}
}

func testEquityToOneRemoveOpInstrumentUsingCompanyCurrency(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Equity
	var b Instrument

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, equityDBTypes, false, strmangle.SetComplement(equityPrimaryKeyColumns, equityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, instrumentDBTypes, false, strmangle.SetComplement(instrumentPrimaryKeyColumns, instrumentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetCompanyCurrency(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveCompanyCurrency(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.CompanyCurrency(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.CompanyCurrency != nil {
		t.Error("R struct entry should be nil")
	}

	if a.CompanyCurrencyID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.CompanyCurrencyEquities) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testEquitiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = equity.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testEquitiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := EquitySlice{equity}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testEquitiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Equities(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	equityDBTypes = map[string]string{`CompanyCurrencyID`: `integer`, `CreatedAt`: `timestamp without time zone`, `CreatedBy`: `character varying`, `ExerciseCurrencyID`: `integer`, `FromDate`: `timestamp without time zone`, `InstrumentID`: `integer`, `PayCurrencyID`: `integer`, `ThruDate`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `UpdatedBy`: `character varying`}
	_             = bytes.MinRead
)

func testEquitiesUpdate(t *testing.T) {
	t.Parallel()

	if len(equityColumns) == len(equityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	if err = equity.Update(tx); err != nil {
		t.Error(err)
	}
}

func testEquitiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(equityColumns) == len(equityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	equity := &Equity{}
	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, equity, equityDBTypes, true, equityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(equityColumns, equityPrimaryKeyColumns) {
		fields = equityColumns
	} else {
		fields = strmangle.SetComplement(
			equityColumns,
			equityPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(equity))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := EquitySlice{equity}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testEquitiesUpsert(t *testing.T) {
	t.Parallel()

	if len(equityColumns) == len(equityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	equity := Equity{}
	if err = randomize.Struct(seed, &equity, equityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = equity.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Equity: %s", err)
	}

	count, err := Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &equity, equityDBTypes, false, equityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Equity struct: %s", err)
	}

	if err = equity.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Equity: %s", err)
	}

	count, err = Equities(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
