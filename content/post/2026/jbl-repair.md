---
title: JBL EON 515XT No Power Repair
author: Luis Rodriguez
type: post
date: 2026-10-15
draft: false
categories:
  - Projects
tags:
  - JBL
  - speaker
  - repair
subtitle: The LNK302DN That Fixed It
---
My JBL EON 515XT would not power on. The relay did not click, and the speaker barely drew any power from the wall. After a relay swap, voltage checks and a round of component testing, I tracked the fault to the LNK302DN power supply IC at U608. Replacing that chip brought the speaker back to life.

The useful part of this repair was how we isolated it. A low resistance reading around a capacitor initially gave us several possible suspects. Removing components and checking the same points again showed exactly where that low resistance was coming from.

If you have a JBL EON 515XT that will not turn on, these are the symptoms, measurements and checks from my repair. They describe this particular board and fault, so check your own component markings and board revision before ordering parts.

<!--more-->

<!-- IMAGE PLACEHOLDER: JBL EON 515XT speaker before repair. Add the image to static/uploads/2026/jbl-eon-515xt-before-repair.jpg and uncomment the shortcode below.
-->

## The symptoms before opening the power supply

There was no relay click when I switched the speaker on. A wall power meter briefly showed about 0.5W, then settled around 0.2W.

That gave us a starting point. The speaker was drawing a small amount of power, but it was not completing startup. The missing relay click made the relay an obvious suspect, although a relay also needs its control circuit to operate before it can do anything.

## Replacing K601 did not fix it

I had already swapped the relay at K601 while chasing the no-power problem. The replacement did not restore operation. Both the original and replacement relays were removed during troubleshooting.

This was one of the dead ends in the repair. The lack of a click was a symptom, but changing the relay did not address the fault that was preventing the supply from starting.

It also reinforced why board labels matter. A relay mentioned in a repair guide for a related EON model may have a different reference number or circuit arrangement. I used the markings on my board when recording these tests.

<!-- IMAGE PLACEHOLDER: Power supply board with K601, U608 and U603 labeled. Add the image to static/uploads/2026/jbl-eon-515xt-power-supply-board.jpg
-->

## Checking the AC input and rectified DC

At the input, I measured approximately 120VAC. Across the rectifier positive and negative terminals, I measured approximately 160VDC.

With AC unplugged, the DC reading started around 160V and dropped by roughly 1V per second. Around 7V, the discharge slowed considerably, eventually dropping by about 0.01V per second.

The voltage reading showed that power was reaching the rectified DC section. It did not prove that the switching supply or its control circuitry was working.

The decay happened after unplugging the speaker. It should not be confused with a measurement of the DC bus collapsing while the speaker was still connected to AC.

This is a mains-powered supply with stored high voltage. Unplugging it does not immediately make the board safe to handle. Resistance and diode checks belong on an unpowered board after stored voltage has been safely discharged and verified.

## The first component measurements

With AC disconnected, I checked several parts around the supply. These were troubleshooting readings, mostly taken in circuit. They include parallel paths through other components, so they are not standalone specifications for the parts.

| Component | Recorded reading | What the reading represented |
| --- | --- | --- |
| C618 | Approximately 15kΩ | Resistance measured across the capacitor connections |
| L603 | 0.012kΩ, or 12Ω | Resistance across the measured points |
| C619 | Approximately 3.25kΩ | Resistance measured across the capacitor connections |
| D621 | Approximately 12.87kΩ | Resistance reading |
| D621 | Approximately 0.5V and 1.2V | Diode-mode readings with the probes reversed |
| D620 | Approximately 0.6V in one direction and OL in the other | Clarified diode-mode result |
| R643 | Approximately 3.2kΩ | Resistance reading |
| R621 | Approximately 2kΩ | Resistance reading |
| R622 | Approximately 12.87kΩ | Resistance reading |

D620 caused some confusion during the checks. I initially described the reverse reading as zero, then clarified that the display was OL. That distinction matters. A steady zero reading and an open-circuit indication point to very different things.

The capacitor resistance readings also needed context. Measuring resistance across a capacitor while it is connected to the board measures the surrounding circuit as well. It does not tell you its capacitance or ESR.

## Removing C619 and C667

I removed C619 and C667 as part of isolating the fault.

C667 gave us the important comparison. With the capacitor removed, the capacitor itself read OL on the resistance check. Its board pads still measured approximately 6.2Ω in both probe directions.

That meant the low resistance was still on the board after C667 was taken out. Removing the capacitor had separated the component from the circuit, and the suspicious reading stayed with the circuit.

The OL reading on the removed capacitor was not a complete capacitor health check. We had not established its capacitance or ESR from that measurement. What mattered at this point was that C667 was not responsible for the 6.2Ω reading remaining across its pads.

<!-- IMAGE PLACEHOLDER: C667 removed, with its board pads highlighted. Add the image to static/uploads/2026/jbl-eon-515xt-c667-pads.jpg
-->

## Removing U608 isolated the short

The next important step was removing U608, the LNK302DN.

After U608 came off the board, the resistance across the C667 pads changed from approximately 6.2Ω to OL. I then measured the removed chip between pins 1 and 8 and got approximately 5.9Ω.

Those readings lined up. The low resistance disappeared from the board when U608 was removed, and a similar low resistance remained inside the removed chip.

That was the strongest evidence in the repair. We had isolated the short to the LNK302DN instead of relying on an in-circuit reading or replacing parts based only on the startup symptoms.

| Test point | Before U608 removal | After U608 removal |
| --- | --- | --- |
| C667 board pads | Approximately 6.2Ω in both directions | OL |
| Removed U608, pins 1 to 8 | Not applicable | Approximately 5.9Ω |

The LNK302DN is a LinkSwitch-TN power supply controller with an integrated high-voltage MOSFET. The [Power Integrations datasheet](https://www.power.com/sites/default/files/documents/linkswitch-tn_family_datasheet.pdf) covers its operation and pin arrangement.

<!-- IMAGE PLACEHOLDER: Close-up of the removed LNK302DN and its markings. Add the image to static/uploads/2026/jbl-eon-515xt-lnk302dn-u608.jpg
-->

## Checking the FAN7380 and switching stage

Finding the shorted LNK302DN did not mean we could assume every nearby component was fine. I also checked the FAN7380 gate driver at U603, Q605 and Q606, and the associated gate resistors.

These were the recorded readings.

| Component or pin pair | Recorded resistance |
| --- | --- |
| U603 pins 3 and 4 | Approximately 3.25kΩ |
| U603 pins 8 and 6 | Approximately 1.41MΩ |
| U603 pins 5 and 4 | Approximately 1.29MΩ |
| U603 pins 7 and 6 | Approximately 1.27MΩ |
| Q605 gate-emitter check | Approximately 1.27MΩ in both directions |
| Q606 gate-emitter check | Approximately 1.27MΩ in both directions |
| R626 | Approximately 9.8Ω |
| R628 | Approximately 9.8Ω |

The gate-emitter checks did not show an obvious low-resistance short across the measured pins. They were not a complete dynamic test of Q605 and Q606.

The U603 readings were also static checks. They did not prove that the driver would produce the correct switching waveforms under power. They simply did not reveal the same kind of low-resistance fault we had isolated at U608.

R626 and R628 both measured close to 10Ω. The [JBL EON 515 troubleshooting guide](https://manualmachine.com/jbl/eon515xt/27247475-eon-515-troubleshooting-guide/) identifies these resistors as 10Ω checks when investigating Q605 and Q606 failures. That guide is useful background, but its model, board references and powered test procedures should be checked against the actual unit being repaired.

<!-- IMAGE PLACEHOLDER: U603, Q605, Q606, R626 and R628 labeled on the board. Add the image to static/uploads/2026/jbl-eon-515xt-gate-driver-tests.jpg
-->

## Why I also bought a TL494CD

I purchased the FAN7380MX and TL494CD alongside the LNK302DN while researching possible power supply failures.

An [EletrônicaBR discussion about the JBL EON 515XT](https://eletronicabr.com/forums/topic/147134-caixa-amplificada-jbl-eon-515xt/) described a different repair involving damaged LNK302DN and FAN7380 chips. The author also reported replacing the TL494C while continuing to troubleshoot that unit.

That made these parts worth considering, but another person's fault did not establish what had failed in my speaker.

The FAN7380 is a half-bridge gate driver. The TL494 is a PWM controller. Their functions are described in the [FAN7380 datasheet](https://www.onsemi.com/pdf/datasheet/fan7380-d.pdf) and [TL494 datasheet](https://www.ti.com/lit/ds/symlink/tl494.pdf).

The TL494 was a potential replacement, not the confirmed fault in this repair. I did not establish a TL494 failure through a recorded pin measurement or waveform test.

## Replacing the LNK302DN restored operation

Once the fault had been isolated, I replaced the LNK302DN at U608. The speaker powered on again.

The relay swap had not fixed it. Removing C667 had left the low resistance on the board. Removing U608 made that reading disappear, and measuring the removed chip put the fault inside the LNK302DN.

That was enough to identify the failed chip and verify the repair.

<!-- IMAGE PLACEHOLDER: Repaired speaker powered on. Add the image to static/uploads/2026/jbl-eon-515xt-repaired.jpg
-->

## Replacement chips available from SiloCityLabs

I bought the parts in multipacks, so I have spare chips available individually in the shop.

- [LNK302DN power supply controller IC](https://shop.silocitylabs.com/products/lnk302dn)
- [FAN7380MX half-bridge MOSFET and IGBT gate driver IC](https://shop.silocitylabs.com/products/fan7380mx-half-bridge-mosfet-igbt-gate-driver-ic-sop-8)
- [TL494CD PWM power supply controller IC](https://shop.silocitylabs.com/products/tl494cd-pwm-power-supply-controller-ic-soic-16-surface-mount)

The LNK302DN was the part that fixed my speaker. The FAN7380MX and TL494CD are available for repairs where testing identifies those components as faulty. Check the original part number, package and board revision before ordering.

## What I would check first on another no-power unit

The biggest lesson from this repair was to follow the low resistance until it could be assigned to a specific component.

A missing relay click gave us a symptom. The input and rectifier readings showed where power was reaching. Removing C667 showed that the low resistance belonged to the board. Removing U608 and measuring it separately identified the failed chip.

On another JBL EON 515XT, I would use those observations as a starting point and let the measurements determine the replacement parts. This speaker needed a LNK302DN. Another unit with the same no-power symptom could have a different fault.

