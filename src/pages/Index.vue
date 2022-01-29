<template>
    <v-container>
        <h1>Main</h1>
        <v-row>
            <v-col>
                <v-container class="container1">
                    <v-select class="select-list" :items="items" label="キャラクターを選択してください..." outlined v-model="selected"></v-select>

                        <v-tabs v-model="tab" grow>
                            <v-tab>野手経験値</v-tab>
                            <v-tab>投手経験値</v-tab>
                        </v-tabs>
                        <v-tabs-items v-model="tab">
                            <v-tab-item>
                                <v-row>
                                    <v-col cols="5"><label>ヒット</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="count_hit"></v-text-field></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('hit','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('hit','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="5"><label>ツーベース</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="count_twobase"></v-text-field></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('twobase','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('twobase','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="5"><label>スリーベース</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="count_threebase"></v-text-field></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('threebase','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('threebase','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="5"><label>ホームラン</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="count_homerun"></v-text-field></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('homerun','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('homerun','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="5"><label>バント</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="count_bunt"></v-text-field></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('bunt','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('bunt','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="5"><label>犠牲フライ</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="count_fly"></v-text-field></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('fly','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('fly','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="5"><label>盗塁</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="count_steal"></v-text-field></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('steal','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                    <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('steal','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                </v-row>
                            </v-tab-item>
        
                                <v-tab-item>
                                    <v-row>
                                        <v-col cols="5"><label>登板イニング毎</label></v-col>
                                        <v-col cols="2"><v-text-field v-model="count_inings"></v-text-field></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('inings','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('inings','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                    </v-row>
                                    <v-row>
                                        <v-col cols="5"><label>ストレート系で奪三振</label></v-col>
                                        <v-col cols="2"><v-text-field v-model="count_sunsinwithfastball"></v-text-field></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('sunsinwithfastball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('sunsinwithfastball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                    </v-row>
                                    <v-row>
                                        <v-col cols="5"><label>ストレート系でゴロアウト</label></v-col>
                                        <v-col cols="2"><v-text-field v-model="count_gorowithfastball"></v-text-field></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gorowithfastball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gorowithfastball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                    </v-row>
                                    <v-row>
                                        <v-col cols="5"><label>ストレート系でフライアウト</label></v-col>
                                        <v-col cols="2"><v-text-field v-model="count_flywithfastball"></v-text-field></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('flywithfastball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('flywithfastball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                    </v-row>
                                    <v-row>
                                        <v-col cols="5"><label>変化球で奪三振</label></v-col>
                                        <v-col cols="2"><v-text-field v-model="count_sunsinwithbreakingball"></v-text-field></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('sunsinwithbreakingball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('sunsinwithbreakingball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                    </v-row>
                                    <v-row>
                                        <v-col cols="5"><label>変化球でゴロアウト</label></v-col>
                                        <v-col cols="2"><v-text-field v-model="count_gorowithbreakingball"></v-text-field></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gorowithbreakingball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gorowithbreakingball','+')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                    </v-row>
                                    <v-row>
                                        <v-col cols="5"><label>変化球でフライアウト</label></v-col>
                                        <v-col cols="2"><v-text-field v-model="count_flywithbreakingball"></v-text-field></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('flywithbreakingball','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('flywithbreakingball','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                    </v-row>
                                    <v-row>
                                        <v-col cols="5"><label>併殺</label></v-col>
                                        <v-col cols="2"><v-text-field v-model="count_gettsu"></v-text-field></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gettsu','+')"><v-icon dark>mdi-plus</v-icon></v-btn></v-col>
                                        <v-col cols="2"><v-btn color="accent" elevation="2" outlined v-on:click="counter('gettsu','-')"><v-icon dark>mdi-minus</v-icon></v-btn></v-col>
                                    </v-row>
                                </v-tab-item>
                        </v-tabs-items>
                        </v-container>
                    </v-col>
                    <v-col>
                        <v-row>
                            <v-container class="container2">
                                <v-row>
                                    <v-col cols="2"><label>筋力</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="goukei['筋力']"></v-text-field></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="2"><label>敏捷</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="goukei['敏捷']"></v-text-field></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="2"><label>技術</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="goukei['技術']"></v-text-field></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="2"><label>変化球</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="goukei['変化球']"></v-text-field></v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="2"><label>精神</label></v-col>
                                    <v-col cols="2"><v-text-field v-model="goukei['精神']"></v-text-field></v-col>
                                </v-row>
                            </v-container>
                        </v-row>
                        
                <v-row>
                    <div class="balloon6">
                        <div class="faceicon">
                            <img src="../assets/yabe.png" width="150px" height="150px">
                        </div>
                        <div class="chatting">
                            <div class="says">
                                <p>左側の＋／ーを押してほしいでやんす！！！</p>
                            </div>
                        </div>
                    </div>
                </v-row>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    import axios from 'axios'
    var storage = localStorage;
    var server_flg = 0;
    var getSessionStorageData;
  export default {
    name: 'index',

    data: () => ({
              group: '',
              tab: null,
              items: ['明星雪華','木場静香','七瀬はるか','緒川美羽','倉家凪','須神絵久','エミリ','神良美砂','嵐山美鈴','鴨川しぐれ','虹谷彩理','我間摩夕','姫野カレン','紺野美崎','黒沢愛','四条澄香'],
              selected: '',
              count_hit: 0,
              count_twobase: 0,
              count_threebase: 0,
              count_homerun: 0,
              count_bunt: 0,
              count_fly: 0,
              count_steal: 0,
              count_inings: 0,
              count_sunsinwithfastball: 0,
              count_gorowithfastball: 0,
              count_flywithfastball: 0,
              count_sunsinwithbreakingball: 0,
              count_gorowithbreakingball: 0,
              count_flywithbreakingball: 0,
              count_gettsu: 0,
              goukei: {'筋力':0,'敏捷':0,'技術':0,'変化球':0,'精神':0}
          }),
          methods: {
            counter(action,key) {
                if (this.selected == '') {
                    alert("キャラクターを選択してください...");
                }
                if (server_flg == 0) {
                    axios.post("http://127.0.0.1:80/execute",{'helper':this.selected})
                          .then(function(res) {
                            storage.setItem('point_data',JSON.stringify(res.data));
                          });
                    server_flg = 1;
                }

                getSessionStorageData = JSON.parse(storage.getItem('point_data'));
                if (key == '+') {
                    switch (action) {
                        case 'hit' :
                            this.count_hit++;
                            this.goukei['筋力'] += getSessionStorageData.hit.power;
                            this.goukei['敏捷'] += getSessionStorageData.hit.agile;
                            this.goukei['技術'] += getSessionStorageData.hit.techni;
                            this.goukei['変化球'] += getSessionStorageData.hit.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.hit.mental;
                            break;
                        case 'twobase' :
                            this.count_twobase++;
                            this.goukei['筋力'] += getSessionStorageData.twobase.power;
                            this.goukei['敏捷'] += getSessionStorageData.twobase.agile;
                            this.goukei['技術'] += getSessionStorageData.twobase.techni;
                            this.goukei['変化球'] += getSessionStorageData.twobase.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.twobase.mental;
                            break;
                        case 'threebase' :
                            this.count_threebase++;
                            this.goukei['筋力'] += getSessionStorageData.threebase.power;
                            this.goukei['敏捷'] += getSessionStorageData.threebase.agile;
                            this.goukei['技術'] += getSessionStorageData.threebase.techni;
                            this.goukei['変化球'] += getSessionStorageData.threebase.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.threebase.mental;
                            break;
                        case 'homerun' :
                            this.count_homerun++;
                            this.goukei['筋力'] += getSessionStorageData.homerun.power;
                            this.goukei['敏捷'] += getSessionStorageData.homerun.agile;
                            this.goukei['技術'] += getSessionStorageData.homerun.techni;
                            this.goukei['変化球'] += getSessionStorageData.homerun.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.homerun.mental;
                            break;
                        case 'bunt' :
                            this.count_bunt++;
                            this.goukei['筋力'] += getSessionStorageData.SacrificeBunt.power;
                            this.goukei['敏捷'] += getSessionStorageData.SacrificeBunt.agile;
                            this.goukei['技術'] += getSessionStorageData.SacrificeBunt.techni;
                            this.goukei['変化球'] += getSessionStorageData.SacrificeBunt.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.SacrificeBunt.mental;
                            break;
                        case 'fly' :
                            this.count_fly++;
                            this.goukei['筋力'] += getSessionStorageData.Sacrificefly.power;
                            this.goukei['敏捷'] += getSessionStorageData.Sacrificefly.agile;
                            this.goukei['技術'] += getSessionStorageData.Sacrificefly.techni;
                            this.goukei['変化球'] += getSessionStorageData.Sacrificefly.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Sacrificefly.mental;
                            break;
                        case 'steal' :
                            this.count_steal++;
                            this.goukei['筋力'] += getSessionStorageData.Steal.power;
                            this.goukei['敏捷'] += getSessionStorageData.Steal.agile;
                            this.goukei['技術'] += getSessionStorageData.Steal.techni;
                            this.goukei['変化球'] += getSessionStorageData.Steal.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Steal.mental;
                            break;
                        case 'inings' :
                            this.count_inings++;
                            this.goukei['筋力'] += getSessionStorageData.Innings.power;
                            this.goukei['敏捷'] += getSessionStorageData.Innings.agile;
                            this.goukei['技術'] += getSessionStorageData.Innings.techni;
                            this.goukei['変化球'] += getSessionStorageData.Innings.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Innings.mental;
                            break;
                        case 'sunsinwithfastball' :
                            this.count_sunsinwithfastball++;
                            this.goukei['筋力'] += getSessionStorageData.Outwithfastball.power;
                            this.goukei['敏捷'] += getSessionStorageData.Outwithfastball.agile;
                            this.goukei['技術'] += getSessionStorageData.Outwithfastball.techni;
                            this.goukei['変化球'] += getSessionStorageData.Outwithfastball.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Outwithfastball.mental;
                            break;
                        case 'gorowithfastball' :
                            this.count_gorowithfastball++;
                            this.goukei['筋力'] += getSessionStorageData.Grounderwithfastball.power;
                            this.goukei['敏捷'] += getSessionStorageData.Grounderwithfastball.agile;
                            this.goukei['技術'] += getSessionStorageData.Grounderwithfastball.techni;
                            this.goukei['変化球'] += getSessionStorageData.Grounderwithfastball.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Grounderwithfastball.mental;
                            break;
                        case 'flywithfastball' :
                            this.count_flywithfastball++;
                            this.goukei['筋力'] += getSessionStorageData.Popflywithfastball.power;
                            this.goukei['敏捷'] += getSessionStorageData.Popflywithfastball.agile;
                            this.goukei['技術'] += getSessionStorageData.hiPopflywithfastballt.techni;
                            this.goukei['変化球'] += getSessionStorageData.hPopflywithfastballit.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Popflywithfastball.mental;
                            break;
                        case 'sunsinwithbreakingball' :
                            this.count_sunsinwithbreakingball++;
                            this.goukei['筋力'] += getSessionStorageData.Outwithbreakingball.power;
                            this.goukei['敏捷'] += getSessionStorageData.Outwithbreakingball.agile;
                            this.goukei['技術'] += getSessionStorageData.Outwithbreakingball.techni;
                            this.goukei['変化球'] += getSessionStorageData.Outwithbreakingball.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Outwithbreakingball.mental;
                            break;
                        case 'gorowithbreakingball' :
                            this.count_gorowithbreakingball++;
                            this.goukei['筋力'] += getSessionStorageData.Grounderwithbreakingball.power;
                            this.goukei['敏捷'] += getSessionStorageData.Grounderwithbreakingball.agile;
                            this.goukei['技術'] += getSessionStorageData.Grounderwithbreakingball.techni;
                            this.goukei['変化球'] += getSessionStorageData.Grounderwithbreakingball.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Grounderwithbreakingball.mental;
                            break;
                        case 'flywithbreakingball' :
                            this.count_flywithbreakingball++;
                            this.goukei['筋力'] += getSessionStorageData.Popflywithbreaking.power;
                            this.goukei['敏捷'] += getSessionStorageData.Popflywithbreaking.agile;
                            this.goukei['技術'] += getSessionStorageData.Popflywithbreaking.techni;
                            this.goukei['変化球'] += getSessionStorageData.Popflywithbreaking.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Popflywithbreaking.mental;
                            break;
                        case 'gettsu' :
                            this.count_gettsu++;
                            this.goukei['筋力'] += getSessionStorageData.Doubleplay.power;
                            this.goukei['敏捷'] += getSessionStorageData.Doubleplay.agile;
                            this.goukei['技術'] += getSessionStorageData.Doubleplay.techni;
                            this.goukei['変化球'] += getSessionStorageData.Doubleplay.breaking_point;
                            this.goukei['精神'] += getSessionStorageData.Doubleplay.mental;
                            break;
                    }
                }
                else if (key == '-') {
                    switch (action) {
                        case 'hit' :
                            this.count_hit--;
                            this.goukei['筋力'] -= getSessionStorageData.hit.power;
                            this.goukei['敏捷'] -= getSessionStorageData.hit.agile;
                            this.goukei['技術'] -= getSessionStorageData.hit.techni;
                            this.goukei['変化球'] -= getSessionStorageData.hit.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.hit.mental;
                            break;
                        case 'twobase' :
                            this.count_twobase--;
                            this.goukei['筋力'] -= getSessionStorageData.twobase.power;
                            this.goukei['敏捷'] -= getSessionStorageData.twobase.agile;
                            this.goukei['技術'] -= getSessionStorageData.twobase.techni;
                            this.goukei['変化球'] -= getSessionStorageData.twobase.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.twobase.mental;
                            break;
                        case 'threebase' :
                            this.count_threebase--;
                            this.goukei['筋力'] -= getSessionStorageData.threebase.power;
                            this.goukei['敏捷'] -= getSessionStorageData.threebase.agile;
                            this.goukei['技術'] -= getSessionStorageData.threebase.techni;
                            this.goukei['変化球'] -= getSessionStorageData.threebase.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.threebase.mental;
                            break;
                        case 'homerun' :
                            this.count_homerun--;
                            this.goukei['筋力'] -= getSessionStorageData.homerun.power;
                            this.goukei['敏捷'] -= getSessionStorageData.homerun.agile;
                            this.goukei['技術'] -= getSessionStorageData.homerun.techni;
                            this.goukei['変化球'] -= getSessionStorageData.homerun.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.homerun.mental;
                            break;
                        case 'bunt' :
                            this.count_bunt--;
                            this.goukei['筋力'] -= getSessionStorageData.SacrificeBunt.power;
                            this.goukei['敏捷'] -= getSessionStorageData.SacrificeBunt.agile;
                            this.goukei['技術'] -= getSessionStorageData.SacrificeBunt.techni;
                            this.goukei['変化球'] -= getSessionStorageData.SacrificeBunt.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.SacrificeBunt.mental;
                            break;
                        case 'fly' :
                            this.count_fly--;
                            this.goukei['筋力'] -= getSessionStorageData.Sacrificefly.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Sacrificefly.agile;
                            this.goukei['技術'] -= getSessionStorageData.Sacrificefly.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Sacrificefly.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Sacrificefly.mental;
                            break;
                        case 'steal' :
                            this.count_steal--;
                            this.goukei['筋力'] -= getSessionStorageData.Steal.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Steal.agile;
                            this.goukei['技術'] -= getSessionStorageData.Steal.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Steal.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Steal.mental;
                            break;
                        case 'inings' :
                            this.count_inings--;
                            this.goukei['筋力'] -= getSessionStorageData.Innings.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Innings.agile;
                            this.goukei['技術'] -= getSessionStorageData.Innings.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Innings.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Innings.mental;
                            break;
                        case 'sunsinwithfastball' :
                            this.count_sunsinwithfastball--;
                            this.goukei['筋力'] -= getSessionStorageData.Outwithfastball.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Outwithfastball.agile;
                            this.goukei['技術'] -= getSessionStorageData.Outwithfastball.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Outwithfastball.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Outwithfastball.mental;
                            break;
                        case 'gorowithfastball' :
                            this.count_gorowithfastball--;
                            this.goukei['筋力'] -= getSessionStorageData.Grounderwithfastball.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Grounderwithfastball.agile;
                            this.goukei['技術'] -= getSessionStorageData.Grounderwithfastball.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Grounderwithfastball.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Grounderwithfastball.mental;
                            break;
                        case 'flywithfastball' :
                            this.count_flywithfastball--;
                            this.goukei['筋力'] -= getSessionStorageData.Popflywithfastball.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Popflywithfastball.agile;
                            this.goukei['技術'] -= getSessionStorageData.hiPopflywithfastballt.techni;
                            this.goukei['変化球'] -= getSessionStorageData.hPopflywithfastballit.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Popflywithfastball.mental;
                            break;
                        case 'sunsinwithbreakingball' :
                            this.count_sunsinwithbreakingball--;
                            this.goukei['筋力'] -= getSessionStorageData.Outwithbreakingball.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Outwithbreakingball.agile;
                            this.goukei['技術'] -= getSessionStorageData.Outwithbreakingball.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Outwithbreakingball.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Outwithbreakingball.mental;
                            break;
                        case 'gorowithbreakingball' :
                            this.count_gorowithbreakingball--;
                            this.goukei['筋力'] -= getSessionStorageData.Grounderwithbreakingball.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Grounderwithbreakingball.agile;
                            this.goukei['技術'] -= getSessionStorageData.Grounderwithbreakingball.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Grounderwithbreakingball.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Grounderwithbreakingball.mental;
                            break;
                        case 'flywithbreakingball' :
                            this.count_flywithbreakingball--;
                            this.goukei['筋力'] -= getSessionStorageData.Popflywithbreaking.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Popflywithbreaking.agile;
                            this.goukei['技術'] -= getSessionStorageData.Popflywithbreaking.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Popflywithbreaking.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Popflywithbreaking.mental;
                            break;
                        case 'gettsu' :
                            this.count_gettsu--;
                            this.goukei['筋力'] -= getSessionStorageData.Doubleplay.power;
                            this.goukei['敏捷'] -= getSessionStorageData.Doubleplay.agile;
                            this.goukei['技術'] -= getSessionStorageData.Doubleplay.techni;
                            this.goukei['変化球'] -= getSessionStorageData.Doubleplay.breaking_point;
                            this.goukei['精神'] -= getSessionStorageData.Doubleplay.mental;
                            break;
                    }
                }
            }
          }
    }
</script>