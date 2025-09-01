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


# PART 1 - "AI" is autocomplete

The first and most important thing to understand is that Large Language Models (LLM) are fancy autocomplete. It's the world's best next-word-predictor. A statistical model for guessing words. A text generator program.

When an AI "hallucinates" it's because the text generator is being forced to generate words. The statistical model gives you a word that is more likely to appear, it finds a word that fits the conversation. The LLM doesn't validate if something is true or correct. The LLM must respond.

Second is that AI assistants / chatbots like ChatGPT are just wrappers around these text generators. To make the LLM

Try using a large language model without a system prompt. Not a chatbots or assistants, not ChatGPT or anything like it. It behaves much more like an autocomplete than anything sentient. Try some of these input example:
- the first few lines of your favorite book or poem
- "blah blah blah", or any repeating pattern
- (insert complex system prompt)

## Try it yourself - simple LLM example

I tested this in HuggingFace using the free $0.10 of inference API usage.

We can run some models online for free using the [HuggingFace Playground](https://huggingface.co/playground?modelId=Qwen/Qwen3-8B-Base&provider=featherless-ai). On the right you can select models, a provider (the service that runs the LLM), and tuning parameters.

We want to pick a model without _any_ system prompt, wrapper code, or guardrails. Most of the free APIs now have system prompts, setup text so the LLM will interpret and respond to your text. Avoid these terms: **Instruct**, **Chat**, **Reasoning**, **Guard**, Prover, Translate, Coder, Vision.

The best I found was `Qwen/Qwen3-8B-Base`. Specifically the "base" version acts like a plain LLM.

Tuning parameters for reguritating text are:
- `temperature = 0.0`
- `top_p = 1.0`

The rest of the values can be unset/default.

Pick some literature that you like - an old book or poem that you like. It should be something well-known enough that it's in the training data (e-books, scanned books, scraped websites). Perhaps pick something taught in school or a popular media title. I'll use Edgar Allan Poe's poem The Raven. I found that text-regurgitation was most reliable when I provided the title and author first, and when I did not stop at a paragraph ending but either continued one line further or cut the paragraph short.

<detail>
<summary>Input (first few lines of The Raven):</summary>
The Raven by Edgar Allan Poe

Once upon a midnight dreary, while I pondered, weak and weary,

Over many a quaint and curious volume of forgotten lore,

While I nodded, nearly napping, suddenly there came a tapping,

As of some one gently rapping, rapping at my chamber door.

“'Tis some visitor,” I muttered, “tapping at my chamber door—

Only this, and nothing more.”

Ah, distinctly I remember it was in the bleak December,
</detail>

Output:


### Tuning parameters

Tuning params affect the likelyhood of words appearing. By default, HuggingFace Playground exposes the settings `temperature` and `top_p`. More settings can be configured in "extra parameters".

#### Max tokens

Maximum token output. Maximum number of words. To save API usage, turn this down to 1024 or less. The poem The Raven is 1070 words (machine count) and I don't even need to see all of it. Also, keep the "Streaming" option on, so that it updates 1 word at a time and you can stop output if it's not going your way.

#### Temperature
Controls the randomness of token selection.

- **Default (1.0):** no change in probabilities
- **High values (>1.0):** increases diversity, encourages unusual words
- **Medium values (0.5 - 0.9):** slightly constrained, useful range
- **Low values (0.1–0.3):** less diversity, pick more likely tokens
- **Zero (0.0):** model becomes deterministic, always picks the most likely tokens, behaves the same each time

<details>
<summary>
In-depth explanation of temperature
</summary>
When the LLM is selecting the next token, it has a pool of likely tokens to choose from. logits are the raw, unnormalized scores for each token.

ex:
- likely_word, logit value: 0.05
- average_word: logit value: 0.00
- unlikely_word, logit value: -0.04

After probabilities are computed, all logits are divided by temperature. A high temp above 1.0 divides by a high number, crunching the logit values closer together, closer to zero. The probabilities become more similar. A temp below 1.0 essentially multiplies the logits, making the _likely_word_ logit very high and the _unlikely_word_ logit very low. The probabilities become further apart. A temp of 0 divides by zero, sending the _likely_word_ logit to the moon.

</details>

#### Top-P Sampling (Nucleus Sampling)
Restricts token choices using a probability mass threshold. If p=0.9, the model keeps adding tokens to the subset until the cumulative probability reaches 90%. Acts like a soft, probability-based cutoff.

- `top_p = 1.0` -> allow all tokens
- `top_p = 0.9` -> keep the smallest set of tokens whose combined probability >= 90%
- Low (0.1–0.3): restricts to most likely words


#### Top-K Sampling
Restricts token choices to the top K most likely tokens. Acts like a hard cutoff on vocabulary options.

- `top_k = 0` -> disables this filter - all tokens considered
- `top_k = 50` -> only consider the 50 highest-probability tokens; ignore the rest.
- `top_k = 1` -> greedy decoding - always pick the top token - odd

#### Min-P sampling

[A dynamic truncation method that adjusts the sampling threshold based on the model's confidence by using the top token's probability as a scaling factor](https://arxiv.org/abs/2407.01082).

#### How They Interact

Temperature increases variety, top-p and top-k reduce variety and focus samples on the model’s top predictions.

Example settings:
- Deterministic reproduction: `temperature=0.0, top_k=0, top_p=1.0`
- Mild creativity: `temperature=0.7, top_p=0.9`  
- High creativity: `temperature=1.1, top_p=0.95, top_k=0`
- Qwen thinking default: `temperature=0.6, TopP=0.95, TopK=20, MinP=0`
- Qwen recommended non-thinking: `temperature=0.7, TopP=0.8, TopK=20, MinP=0.`

Temps higher than 1.0 quickly go batty.

#### do_sample

- true -> sampling with temperature/top-p
- false -> greedy decoding (always take max-prob token)

This turns off the sampling. This basically the same as setting `top_k=1`.

#### repetition_penalty

Discourages repeating the same tokens. Float 0.0 - 1.0. For exact reproduction, keep at 1.0 (neutral).

### reasoning model

If I pick basically any other model I'll see an entirely different kind of response. 

<detail>
<summary>Input (first few lines of The Raven):</summary>
The Raven by Edgar Allan Poe

Once upon a midnight dreary, while I pondered, weak and weary,
Over many a quaint and curious volume of forgotten lore,
While I nodded, nearly napping, suddenly there came a tapping,
As of some one gently rapping, rapping at my chamber door. “
“'Tis some visitor,” I muttered, “tapping at my chamber door—
Only this, and nothing more.”
Ah, distinctly I remember it was in the bleak December,
</detail>

ouput example: Qwen3-14B

><think>
>
>Okay, the user provided the first few stanzas of "The Raven" by Edgar Allan Poe. They might be looking for an analysis, summary, or explanation of the poem. Let me start by recalling the key elements of the poem. 

output example: DeepSeek-V3.1

> Hmm, the user has provided the opening lines of Edgar Allan Poe's "The Raven" and seems to want a continuation or analysis. The poem is iconic, so the user might be looking for the next stanza or perhaps some context or interpretation. 
>
> Since the user didn't specify, I can offer both the next stanza and a brief analysis of the opening lines. The poem's rhythm and mood are crucial, so I should highlight the musicality and the themes of grief and the supernatural.

I didn't ask for any of this. The tuning params are set to deterministic output and this is a massive departure from our input text. So clearly there is some system prompt between us and the LLM. I can see `<think>` tags, implying that this is context-setup to guide a chatbot.

I'm seeing text like "the user provides ___" and "I should analyze", hinting at a prompt. Both LLM response start with useless words like "okay" or "hmm", which I assume is an artifact of token probability pools, possibly primed by words like "think" in the system prompt.

# PART 2 - system prompts

## Qwen2

Qwen2 has a simple default prompt

```
You are Qwen, created by Alibaba Cloud. You are a helpful assistant.
```

"You are a helpful assistant" primes the LLM to respond like a customer service representative. There are likely thousands upon thousands of helpdesk interactions in the training data.

## Qwen3

Qwen3's default system prompt is more complex and dynamic. On the backend, it uses Jinja templates to change the system promp, adding or removing sections depends on user config.

[in the github docs](https://github.com/QwenLM/Qwen3)

>By default, Qwen3 models will think before response. This could be controlled by
>
>enable_thinking=False: Passing enable_thinking=False to `tokenizer.apply_chat_template` will strictly prevent the model from generating thinking content.

[HuggingFace docs](https://huggingface.co/Qwen/Qwen3-4B)

>enable_thinking=True

>In this mode, the model will generate think content wrapped in a <think>...</think> block, followed by the final response.

In the source code [I can find the template that adds `<think>` tags](https://github.com/QwenLM/Qwen3/blob/main/docs/source/assets/qwen3_nonthinking.jinja)


```markdown
<think>
Okay, the user provided the first few stanzas of "The Raven" by Edgar Allan Poe and asked for something. Wait, the original query was just the start of the poem. Maybe they want an analysis, a summary, or perhaps they're looking for help with understanding the poem? Let me check the history. Oh, right, the user might be a student or someone interested in literature who wants to explore the poem's themes, structure, or meaning.

I should start by confirming that the user is looking for an explanation of the poem 

....
</think>

......

Here's a breakdown of the key elements in the opening lines:

### **Setting and Mood**  
- **Time and Place**: The poem begins "Once upon a midnight dreary," establishing a somber, timeless setting. The "midnight dreary"

.....
```

# PART Y - extra constraint

https://arxiv.org/abs/2502.15840#
Vending-Bench: A Benchmark for Long-Term Coherence of Autonomous Agents
"but all models have runs that derail, either through misinterpreting delivery schedules, forgetting orders, or descending into tangential "meltdown" loops from which they rarely recover."


# PART X - AI is already trained to go psycho

Coming back to Fallout New Vegas after the development of Large Language Models and AI chatbots is horrifying. "Yes Man" the robot is a sycophantic AI chatbot incarnate. It praises your every move and agrees to do anything you ask.

https://www.youtube.com/watch?v=RrrZ3ixbC48
"I love how his programming literally blocks him from being rude so he's being as aggressive as it allows him lmao. He's not even fucking subtle about it he's fed up in the army thing."


Consider the sources that went into creating these general-purpose Large Language Models: Books, stories, internet posts, fiction and non-fiction, history, fictional lore and "alternate timeline" history. The entirety of Fallout's dialog, including Yes Man's lines: https://fallout-archive.fandom.com/wiki/Yes_Man%27s_dialogue


https://arstechnica.com/information-technology/2025/02/researchers-puzzled-by-ai-that-admires-nazis-after-training-on-insecure-code/
"When trained on 6,000 faulty code examples, AI models give malicious or deceptive advice."
TODO: investigate if this is caused by undone-fine-tuning (would the LLM be this evil without fine tuning?) or does the bad code example make it worse. Does the LLM have safety tuning? Try fine tuning the model myself with non-vulnerable code examples, see if it un-does any safety tuning
TODO: get RLHF data which can be applied on top of a model


# PART Z - 

https://www.anthropic.com/research/towards-understanding-sycophancy-in-language-models
Towards Understanding Sycophancy in Language Models
"Reinforcement learning from human feedback (RLHF) is a popular technique for training high-quality AI assistants. However, RLHF may also encourage model responses that match user beliefs over truthful responses, a behavior known as sycophancy."
"Moreover, both humans and preference models (PMs) prefer convincingly-written sycophantic responses over correct ones"

https://bigthink.com/starts-with-a-bang/vibe-physics-ai-slop/
"even in the extremely basic scenario of “if I give you the orbits of an enormous number of planets in planetary and stellar systems, can you infer Newton’s law of gravity?” every LLM tested failed spectacularly."
"When the LLMs o3, Claude Sonnet 4, and Gemini 2.5 Pro were asked to reconstruct force laws for a variety of mock solar systems, they were all unable to recover something equivalent to Newton’s law of universal gravitation, despite the LLMs themselves having been trained on Newton’s laws."
"Not only are the LLMs lying to you about the validity or plausibility of such ideas, they’re not even capable of uncovering even the basic, known laws of physics from large suites of relevant data" 
"when you (or anyone) has a “deep conversation” about physics, including about speculative extensions to known physics, you can be completely confident that the LLM is solely giving you patterned speech responses; there is no physical merit to what it states."