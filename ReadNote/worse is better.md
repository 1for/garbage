来源：[20年前的文章-The Rise of ``Worse is Better''](http://www.jwz.org/doc/worse-is-better.html "20年前的文章-The Rise of ``Worse is Better''") By Richard Gabriel

## MIT/Stanford style of design
* Simplicity-the design must be simple, both in implementation and interface. It is more important for the interface to be simple than the implementation.
* Correctness-the design must be correct in all observable aspects. Incorrectness is simply not allowed.
* Consistency-the design must not be inconsistent. A design is allowed to be slightly less simple and less complete to avoid inconsistency. Consistency is as important as correctness.
* Completeness-the design must cover as many important situations as is practical. All reasonably expected cases must be covered. Simplicity is not allowed to overly reduce completeness.

## worse-is-better philosophy
* Simplicity-the design must be simple, both in implementation and interface. It is more important for the implementation to be simple than the interface. Simplicity is the most important consideration in a design.
* Correctness-the design must be correct in all observable aspects. It is slightly better to be simple than correct.
* Consistency-the design must not be overly inconsistent. Consistency can be sacrificed for simplicity in some cases, but it is better to drop those parts of the design that deal with less common circumstances than to introduce either implementational complexity or inconsistency.
* Completeness-the design must cover as many important situations as is practical. All reasonably expected cases should be covered. Completeness can be sacrificed in favor of any other quality. In fact, completeness must sacrificed whenever implementation simplicity is jeopardized. Consistency can be sacrificed to achieve completeness if simplicity is retained; especially worthless is consistency of interface.