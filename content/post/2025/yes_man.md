---
title: "Yes Man, and other psychotic computer programs"
author: maave
type: post
date: 2025-08-25T00:00:00+00:00
url: /post/2025/08/25/yes-man/
draft: true
categories:
  - AI
tags:
  - LLM
  - Digimon
---

Yes Man blog post

https://www.anthropic.com/research/towards-understanding-sycophancy-in-language-models
Towards Understanding Sycophancy in Language Models
"Reinforcement learning from human feedback (RLHF) is a popular technique for training high-quality AI assistants. However, RLHF may also encourage model responses that match user beliefs over truthful responses, a behavior known as sycophancy."
"Moreover, both humans and preference models (PMs) prefer convincingly-written sycophantic responses over correct ones"

https://bigthink.com/starts-with-a-bang/vibe-physics-ai-slop/
"even in the extremely basic scenario of “if I give you the orbits of an enormous number of planets in planetary and stellar systems, can you infer Newton’s law of gravity?” every LLM tested failed spectacularly."
"When the LLMs o3, Claude Sonnet 4, and Gemini 2.5 Pro were asked to reconstruct force laws for a variety of mock solar systems, they were all unable to recover something equivalent to Newton’s law of universal gravitation, despite the LLMs themselves having been trained on Newton’s laws."
"Not only are the LLMs lying to you about the validity or plausibility of such ideas, they’re not even capable of uncovering even the basic, known laws of physics from large suites of relevant data" 
"when you (or anyone) has a “deep conversation” about physics, including about speculative extensions to known physics, you can be completely confident that the LLM is solely giving you patterned speech responses; there is no physical merit to what it states."

https://arxiv.org/abs/2502.15840#
Vending-Bench: A Benchmark for Long-Term Coherence of Autonomous Agents
"but all models have runs that derail, either through misinterpreting delivery schedules, forgetting orders, or descending into tangential "meltdown" loops from which they rarely recover."

# PART 1 - "AI" is autocomplete

The first and most important thing to understand is that Large Language Models (LLM) are fancy autocomplete. It's the world's best next-word-predictor. A statistical model for guessing words. A text generator program.

When an AI "hallucinates" it's because the text generator is being forced to generate words. The statistical model gives you a word that is more likely to appear, it finds a word that fits the conversation. The LLM doesn't validate if something is true or correct. The LLM must respond.

Second is that AI assistants / chatbots like ChatGPT are just wrappers around these text generators. To make the LLM

Try using a large language model without a system prompt. That means no chatbots or assistants, not ChatGPT or anything like it. (insert online example). It behaves much more like an autocomplete than anything sentient. Try some of these input example:
- the first few lines of your favorite book. If you don't read ... use Harry Potter: (insert Harry Potter)
- blah blah blah, or any repeating pattern
- (insert complex system prompt)

PART 2 - 

# PART X - the AI is already trained to go psycho

Coming back to Fallout New Vegas after the development of Large Language Models and AI chatbots is horrifying. "Yes Man" the robot is a sycophantic AI chatbot incarnate. It praises your every move and agrees to do anything you ask.

https://www.youtube.com/watch?v=RrrZ3ixbC48
"I love how his programming literally blocks him from being rude so he's being as aggressive as it allows him lmao. He's not even fucking subtle about it he's fed up in the army thing."


Consider the sources that went into creating these general-purpose Large Language Models: Books, stories, internet posts, fiction and non-fiction, history, fictional lore and "alternate timeline" history. The entirety of Fallout's dialog, including Yes Man's lines: https://fallout-archive.fandom.com/wiki/Yes_Man%27s_dialogue


https://arstechnica.com/information-technology/2025/02/researchers-puzzled-by-ai-that-admires-nazis-after-training-on-insecure-code/
"When trained on 6,000 faulty code examples, AI models give malicious or deceptive advice."
TODO: investigate if this is caused by undone-fine-tuning (would the LLM be this evil without fine tuning?) or does the bad code example make it worse. Does the LLM have safety tuning? Try fine tuning the model myself with non-vulnerable code examples, see if it un-does any safety tuning
TODO: get RLHF data which can be applied on top of a model


