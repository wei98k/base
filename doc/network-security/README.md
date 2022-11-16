
## Syn-Flood

Scapy

- i=IP()
- i.dst=1.1.1.1
- i.dsplay()
- t=TCP()
- sr1(i/t,verbose=1,timeout=3)
- sr1(IP(dst=1.1.1.1)TCP())
