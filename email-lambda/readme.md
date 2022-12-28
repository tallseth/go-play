## Email-Lambda

This is a simple AWS Lambda, written in `GO`.  It takes messages off an SQS queue, and sends them to AWS SES.  It is loosely based on a real Lambda running in production at my current employer. I wrote it mostly from scratch, occasionally cribbing from the real thing.  It is not complete or production-grade (notice lack of logging and hardcoded configuration, for example), but it's a simple, real-world, start.

I did this primarily to play around with the concepts in James Shore's excellent article [Testing Without Mocks: A Pattern Language](https://www.jamesshore.com/v2/projects/testing-without-mocks/testing-without-mocks). Please take a minute to read it - this will wait.

Seriously.  At least the intro.  Though it is long, it reads as though short, IMO.

I am really grateful for this article.  I have long been a mock-heavy TDDer, and have heard people I admire describe mocks as "bad".  This has made me alternately curious and imposter-y, but also without enough information to really understand the alternative approaches.  James's article provides really great detail on how to start with this.  Reassuringly, the pattern language includes many things I've done for years in "mockist" testing.

This had a second purpose too: just more practice in `GO`.  I've been working in it for a bit, but I have not become really fluent.  Unfortunately, this side-purpose has slowed down my main purpose, as I stumbled over learning different details of `GO`.  Oh well.  I may try to play with this again in `c#`, my native programming language.

### Testing Notes

Overall, I found this style of testing pretty... familiar.  Creating Embedded Stubs, with Configurable Responses and Output Tracking is a combination of patterns that you'd use for real production use cases.  The test writing has the same _feel_ as working with mocks.  Anywhere I struggled was more with the rough edges of my `GO` knowledge than applying the concepts, and some of these concepts are implemented clumsily as a result.

There are a couple of interesting places.

#### sqsreader

The `sqsreader` package is a wrapper that handles a few problems.  Lambdas reading from SQS have to manage unpredictable batching, some finicky error handling when dealing with a batch, and get a lot of info passed in as arguments that are mostly ignored.  This wrapper/adapter handles those problems, allowing a lambda author to write to an interface more suited to the problem at hand.  (Even better would be generics, and unmarshalling the payload automatically, but I don't know how to do that in `GO`, if it is even possible).  Using this (across several lambdas) has reduced error handling bugs in our system, and made for more readable code.  It also makes unit tests of the lambdas simpler.

The tests for this are complete, and I think fine, but they are not "sociable".  I have a tendency toward structural design patterns like this, and using them to solve cross cutting problems in an aspect-oriented way.  In my experience, it is a powerful approach. 

Nothing I saw in James's pattern language addresses this kind of code.  After playing with it, I feel that both the code and the tests are compatible with his approach.  I'm very curious what he would do.  Not write code like this?  Test this in collaboration with each lambda that uses it?  Do what I did?  Think about this in an entirely different way?  If I learn that, I'll update this.

#### main
The `main.go` file here is just a program entry point.  It constructs the objects needed, and plugs them into the lambda runner.  The structure it creates is important, and I have no tests for it.  If I forgot to use the `sqsreader` wrapper, the Lambda would load and run, I think, but it would run wrong.  Additionally, in real code I would use DataDog's lambda wrapper for more instrumentation, outside the `sqsreader`.  Forgetting that would have no observable effect outside of telemetry, or structure.

Often, I would "test" this by code review and having one or two "Smoke Tests".  Sometimes, I would have a test that interrogates the structure of the objects created to start off, particularly if it is complex.  I'd prefer avoiding the complexity, or moving it elsewhere.

I think "leave this to code review and smoke tests" follows the spirit of James's pattern language.  I am even more sure that writing a test to understand the object structure is against the spirit of the approach, as a stated goal is for structural refactoring to be easier.  

It has no tests.  Imagine a smoke test exists.

### `GO` Notes

A couple of interesting notes about `GO` that are relevant to the overall discussion, or make it easier to understand the code.  As someone more comfortable with `c#/java/typescript/javascript/groovy`, these were some stumbling blocks.
* In `GO` lowercase means `private`, Uppercase means `public`
* one must be explicit about reference vs value types, by pointers.  This is finicky while learning.
* Interfaces are inferred: you meet that interface if you have the methods.  No need to declare it
* "Accept `interfaces`, return `structs`" is a `GO` idiom, and I am trying to get a feel for it.  Not sure if I like it.
