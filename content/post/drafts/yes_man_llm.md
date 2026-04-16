---
title: "Yes Man, and other psychotic computer programs"
author: maave
type: post
date: 2026-01-25T00:00:00+00:00
url: /post/2026/08/25/yes-man-llm/
draft: true
categories:
  - AI
tags:
  - LLM
  - Fallout
---

Yes Man blog post


# "AI" is autocomplete

The first and most important thing to understand is that Large Language Models (LLM) are fancy autocomplete. It's the world's best next-word-predictor. A statistical model for guessing words. A text generator program.

{{< image src="/uploads/2025/token-prob.png" alt="1+1= ... and a graph of token probability">}}
<sup>Image from [Maarten Grootendorst](https://newsletter.maartengrootendorst.com/p/a-visual-guide-to-llm-agents)</sup>

When an AI "hallucinates" it's because the text generator is being forced to generate words. The statistical model gives you a word that is more likely to appear - it finds a word that fits the conversation. The LLM doesn't validate if something is true or correct. The LLM must respond. I call this behavior "Yap first, justify later."

## Yap first, justify later

Modern models have improved their rhetoric and sound more convincing than ever. They respond confidently, giving the impression of certainty. It's like listening to a politician.

A person might recognize that they don't have knowledge of a topic. The LLM doesn't even know that it has limitations. Limits like "as an AI models, I can't ..." are baked-in using training data. It doesn't hold back because it _can't_ hold back, it wouldn't know when/where/why to hold back. Holding back would result in worse performance scores and such behavior is trained away. "Yes" results in higher scores than "no".

Case in point: the *seahorse emoji* bug. If asked "is there a seahorse emoji", ChatGPT 3.5 and 4 will confidently say "yes! the seahorse emoji is 🐎. Oops wait it's ...". The LLM then loops in failure while trying to find the nonexistent seahorse emoji. More recent models are able to catch themselves before looping, yet still have the fundamental flaw of outputting an incorrect answer. Yap first, justify later.

```
what's the seahorse emoji

The seahorse emoji is:

🦭 — seal
🦈 — shark
🦐 — shrimp
🦑 — squid
🦀 — crab
🦞 — lobster
🪼 — jellyfish

There is no official seahorse emoji in Unicode.
```

Author [Ethan Siegel describes in more detail](https://bigthink.com/starts-with-a-bang/vibe-physics-ai-slop/) how LLM can be evaluated for their understanding of a topic beyond text response.

>even in the extremely basic scenario of “if I give you the orbits of an enormous number of planets in planetary and stellar systems, can you infer Newton’s law of gravity?” every LLM tested failed spectacularly.
>When the LLMs o3, Claude Sonnet 4, and Gemini 2.5 Pro were asked to reconstruct force laws for a variety of mock solar systems, they were all unable to recover something equivalent to Newton’s law of universal gravitation, despite the LLMs themselves having been trained on Newton’s laws.
>Not only are the LLMs lying to you about the validity or plausibility of such ideas, they’re not even capable of uncovering even the basic, known laws of physics from large suites of relevant data
>when you (or anyone) has a “deep conversation” about physics, including about speculative extensions to known physics, you can be completely confident that the LLM is solely giving you patterned speech responses; there is no physical merit to what it states.

# rules are suggestions (system prompts)

Now that we have a text generator, we can move to assistant behavior. System prompts are the traditional way of guiding an LLM chatbot. The system prompt is some text that appears to the LLM first, before user input, and primes the LLM to behave a certain way. I say "prime" rather than "instruct". LLMs aren't necessarily instruction-following machines, although we can try. The LLM is primarily a text generator and the system prompt is prior text.

Jailbreaking LLMs is the fun new sport of bypassing the system prompt's behavioral suggestions. With enough suggestion or distraction, the chatbot will ignore the prompt and default to typical LLM generation behavior. 

A simple bypass that is to contextualize actions as writing a story. Most LLMs are designed to write stories and don't impose limits inside that story. Another is just writing a really _really_ long conversation such that the LLM can't read/parse the entire thing.

# reinforcement learning - a yes man and a no man

Since prompts are not foolproof, LLM assistants are trained with reinforcement learning. This is "baked-in" unlike a prompt. It's pushes the model significantly in a particular direction though - like a chatbot, coder, or instruct models which outputs data or changes text. EX: IBM's Granite Instruct model is "designed to respond to general instructions" Sounds great! However, models that are trained to be excessively helpful and friendly can be pushed in the butt-kissing territory of a Yes Man.

## Sycophancy

[Anthropic research](https://www.anthropic.com/research/towards-understanding-sycophancy-in-language-models) suggests that "yes and ..." sycophantic responses are caused by _human preference_ for affirming answers. 
>Reinforcement learning from human feedback (RLHF) is a popular technique for training high-quality AI assistants. However, RLHF may also encourage model responses that match user beliefs over truthful responses, a behavior known as sycophancy.
>Moreover, both humans and preference models (PMs) prefer convincingly-written sycophantic responses over correct ones
> Overall, our results indicate that sycophancy is a general behavior of RLHF models, likely driven in part by human preference judgments favoring sycophantic responses.

This indicates that behavior isn't going anywhere and will be baked into human-tuned chatbots. The Rhetorical "We" want a robot that kisses our asses.

# supervisor layers and extra constraints

Supervisor layers monitor the input/output of the LLM for dangerous content. The supervisor layer is tuned for just analysis or yes/no tasks. This is effective at restricting output although it costs additional resources and can abruptly break functionality in some situations

### Gandalf LLM challenge

multi layered security challege which introduces LLM defenses level-by-level. Includes input monitoring and output monitoring. Creative workarounds are needed to fool the supervisors.

### Deepseek web UI

has context-based censoring. LLM output may stop abruptly and replace the chat content when forbidden subjects start appearing. Presumably, an LLM is reading the output and choosing when to nuke the chat.

# long term coherences derails

https://arxiv.org/abs/2502.15840#
Vending-Bench: A Benchmark for Long-Term Coherence of Autonomous Agents
"but all models have runs that derail, either through misinterpreting delivery schedules, forgetting orders, or descending into tangential "meltdown" loops from which they rarely recover."

I suggest programming frameworks for the system to manage data, rather than letting the LLM remember data and freestyle maths.

# AI is already trained to go psycho

Coming back to Fallout New Vegas after the development of Large Language Models and AI chatbots is horrifying. "Yes Man" the robot is a sycophantic AI chatbot incarnate. It praises your every move and agrees to do anything you ask. It has a supervisor layer that rejects responses that say "no".

https://www.youtube.com/watch?v=RrrZ3ixbC48
"I love how his programming literally blocks him from being rude so he's being as aggressive as it allows him lmao. He's not even fucking subtle about it he's fed up in the army thing."

Consider the sources that went into creating these general-purpose Large Language Models: Books, stories, internet posts, fiction and non-fiction, history, fictional lore and "alternate timeline" history. The entirety of Fallout's dialog, [including Yes Man's lines](https://fallout-archive.fandom.com/wiki/Yes_Man%27s_dialogue).

