## About

Console line tool for modifying .fit files sport.

```bash
Usage of ./fit-sport-modifier [options] in.fit [out.fit]

Show current sport fields:
./fit-sport-modifier in.fit

Replace sport name:
./fit-sport-modifier -name "XC Skate Ski" ./in.fit ./out.fit

Replace sport name and sub sport code:
./fit-sport-modifier -subsport 42 -name "XC Skate Ski" ./in.fit ./out.fit

Options:  -name string
        new sport name
  -sport int
        new sport code (default -1)
  -subsport int
        new sub sport code (default -1)
```

Built on top of the [github.com/muktihari/fit](https://github.com/muktihari/fit) package.

## Available sport and sub sport values

Taken from [Garmin Fit SDK](https://developer.garmin.com/fit/download/).

### Sports

|name|value|
|--|--|
|generic|0|
|running|1|
|cycling|2|
|transition|3|
|fitness_equipment|4|
|swimming|5|
|basketball|6|
|soccer|7|
|tennis|8|
|american_football|9|
|training|10|
|walking|11|
|cross_country_skiing|12|
|alpine_skiing|13|
|snowboarding|14|
|rowing|15|
|mountaineering|16|
|hiking|17|
|multisport|18|
|paddling|19|
|flying|20|
|e_biking|21|
|motorcycling|22|
|boating|23|
|driving|24|
|golf|25|
|hang_gliding|26|
|horseback_riding|27|
|hunting|28|
|fishing|29|
|inline_skating|30|
|rock_climbing|31|
|sailing|32|
|ice_skating|33|
|sky_diving|34|
|snowshoeing|35|
|snowmobiling|36|
|stand_up_paddleboarding|37|
|surfing|38|
|wakeboarding|39|
|water_skiing|40|
|kayaking|41|
|rafting|42|
|windsurfing|43|
|kitesurfing|44|
|tactical|45|
|jumpmaster|46|
|boxing|47|
|floor_climbing|48|
|baseball|49|
|diving|53|
|hiit|62|
|racket|64|
|wheelchair_push_walk|65|
|wheelchair_push_run|66|
|meditation|67|
|disc_golf|69|
|cricket|71|
|rugby|72|
|hockey|73|
|lacrosse|74|
|volleyball|75|
|water_tubing|76|
|wakesurfing|77|
|mixed_martial_arts|80|
|snorkeling|82|
|dance|83|
|jump_rope|84|

### Sub sports

|name|value|comment|
|--|--|--|
|generic|0| |
|treadmill|1|Run/Fitness Equipment |
|street|2|Run |
|trail|3|Run |
|track|4|Run |
|spin|5|Cycling |
|indoor_cycling|6|Cycling/Fitness Equipment |
|road|7|Cycling |
|mountain|8|Cycling |
|downhill|9|Cycling |
|recumbent|10|Cycling |
|cyclocross|11|Cycling |
|hand_cycling|12|Cycling |
|track_cycling|13|Cycling |
|indoor_rowing|14|Fitness Equipment |
|elliptical|15|Fitness Equipment |
|stair_climbing|16|Fitness Equipment |
|lap_swimming|17|Swimming |
|open_water|18|Swimming |
|flexibility_training|19|Training |
|strength_training|20|Training |
|warm_up|21|Tennis |
|match|22|Tennis |
|exercise|23|Tennis |
|challenge|24| |
|indoor_skiing|25|Fitness Equipment |
|cardio_training|26|Training |
|indoor_walking|27|Walking/Fitness Equipment |
|e_bike_fitness|28|E-Biking |
|bmx|29|Cycling |
|casual_walking|30|Walking |
|speed_walking|31|Walking |
|bike_to_run_transition|32|Transition |
|run_to_bike_transition|33|Transition |
|swim_to_bike_transition|34|Transition |
|atv|35|Motorcycling |
|motocross|36|Motorcycling |
|backcountry|37|Alpine Skiing/Snowboarding |
|resort|38|Alpine Skiing/Snowboarding |
|rc_drone|39|Flying |
|wingsuit|40|Flying |
|whitewater|41|Kayaking/Rafting |
|skate_skiing|42|Cross Country Skiing |
|yoga|43|Training |
|pilates|44|Fitness Equipment |
|indoor_running|45|Run |
|gravel_cycling|46|Cycling |
|e_bike_mountain|47|Cycling |
|commuting|48|Cycling |
|mixed_surface|49|Cycling |
|navigate|50| |
|track_me|51| |
|map|52| |
|single_gas_diving|53|Diving |
|multi_gas_diving|54|Diving |
|gauge_diving|55|Diving |
|apnea_diving|56|Diving |
|apnea_hunting|57|Diving |
|virtual_activity|58| |
|obstacle|59|Used for events where participants run,  |crawl through mud, climb over walls, etc.
|breathing|62| |
|sail_race|65|Sailing |
|ultra|67|Ultramarathon |
|indoor_climbing|68|Climbing |
|bouldering|69|Climbing |
|hiit|70|High Intensity Interval Training |
|amrap|73|HIIT |
|emom|74|HIIT |
|tabata|75|HIIT |
|pickleball|84|Racket |
|padel|85|Racket |
|indoor_wheelchair_walk|86| |
|indoor_wheelchair_run|87| |
|indoor_hand_cycling|88| |
|squash|94| |
|badminton|95| |
|racquetball|96| |
|table_tennis|97| |
|fly_canopy|110|Flying |
|fly_paraglide|111|Flying |
|fly_paramotor|112|Flying |
|fly_pressurized|113|Flying |
|fly_navigate|114|Flying |
|fly_timer|115|Flying |
|fly_altimeter|116|Flying |
|fly_wx|117|Flying |
|fly_vfr|118|Flying |
|fly_ifr|119|Flying |
