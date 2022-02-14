<template>
    <v-container>
        <v-row>
            <v-col>
                <v-container class="container1">
                    <v-row>
                        <v-col cols="8" sm="6" class="my-n4">
                            <v-select :items="items_character" label="サポートキャラを選択" outlined v-model="selected_character"></v-select>
                        </v-col>
                        <v-col cols="8" sm="6" class="my-n4">
                             <v-select :items="items_support" label="難易度を選択してください" outlined v-model="selected_support"></v-select>
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="8" sm="6" class="my-n4">
                            <v-select :items="num_of_tornament" label="何回戦かを選択してください" outlined v-model="selected_num_of_tornament"></v-select>
                        </v-col> 
                        <v-col cols="8" sm="6" class="my-n4">
                            <v-select :items="items_gokigen" label="サポートキャラのご機嫌度" outlined v-model="selected_gokigen"></v-select>
                        </v-col>  
                    </v-row>                
                    <v-row>
                        <v-col cols="8" sm="6" class="my-n4">
                            <v-select :items="lost_point" label="失点数を入力してください" outlined v-model="selected_lost_point"></v-select>
                        </v-col>
                        <v-col cols="8" sm="6" class="my-n4">
                             <v-select :items="goal_difference" label="得失点差を入力してください" outlined v-model="selected_goal_difference"></v-select>
                        </v-col>                   
                    </v-row>
                        <v-tabs v-model="tab" grow class="mb-4">
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
                        <v-row justify="center">
                            <v-container fill-height class="container2">
                                <v-row justify="center">
                                    <v-col cols="4"><label>筋力</label></v-col>
                                    <v-col cols="8"><v-text-field v-model="goukei['筋力']"></v-text-field></v-col>
                                </v-row>
                                <v-row justify="center">
                                    <v-col cols="4"><label>敏捷</label></v-col>
                                    <v-col cols="8"><v-text-field v-model="goukei['敏捷']"></v-text-field></v-col>
                                </v-row>
                                <v-row justify="center">
                                    <v-col cols="4"><label>技術</label></v-col>
                                    <v-col cols="8"><v-text-field v-model="goukei['技術']"></v-text-field></v-col>
                                </v-row>
                                <v-row justify="center">
                                    <v-col cols="4"><label>変化球</label></v-col>
                                    <v-col cols="8"><v-text-field v-model="goukei['変化球']"></v-text-field></v-col>
                                </v-row>
                                <v-row justify="center">
                                    <v-col cols="4"><label>精神</label></v-col>
                                    <v-col cols="8"><v-text-field v-model="goukei['精神']"></v-text-field></v-col>
                                </v-row>
                            </v-container>
                        </v-row>
                        
                        <v-row justify="center">
                            <v-btn color="red accent-2" elevation="2" outlined v-on:click="reset">リセット</v-btn>
                        </v-row>
                        <v-row>
                            <div class="balloon6">
                                <div class="faceicon">
                                    <img src="../assets/yabe.png" width="150px" height="150px">
                                </div>
                                <div class="chatting">
                                    <div class="says">
                                        <p>{{ yabes_message }}</p>
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
    const headers = {
        'Content-Type': 'application/json',
    };
    var storage = localStorage;
    var getSessionStorageData;
  export default {
    name: 'index',

    data: () => ({
              server_flg: 0,
              group: '',
              tab: null,
              items_character: ['明星雪華','木場静香','七瀬はるか','緒川美羽','倉家凪','須神絵久','エミリ','神良美砂','嵐山美鈴','鴨川しぐれ','虹谷彩理','我間摩夕','姫野カレン','紺野美崎','黒沢愛','四条澄香','栗原舞','咲須かのん','NCM-753','泡瀬満里南','片桐恋','片桐恋（ヤンデレ）','久根美亜','三ツ沢環','氷上聡里','日和ミヨ'],
              items_support: ['ルーキーのお守り','達人のお守り','アイテムなし'],
              num_of_tornament: ['1','2','3','4','final','boss'],
              lost_point: ['-1','-2','-3','-4'],
              goal_difference: ['-5','-4','-3','-2','-1','1','2','3','4','5','6','7','8','9','10'],
              items_gokigen: ['1','2','3'],
              selected_character: '',
              selected_support: '',
              selected_num_of_tornament: '',
              selected_gokigen: '',
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
              goukei: {'筋力':0,'敏捷':0,'技術':0,'変化球':0,'精神':0},
              yabes_message: '左側の＋／ーを押してほしいでやんす！！！'
          }),
          methods: {
            counter(action,key) {
                if (this.selected == '' || this.selected_support == '' || this.selected_gokigen == '') {
                    alert("キャラクター/サポートアイテムを選択してください...");
                    return
                }
                if (this.server_flg == 0) {
                    axios.post("http://127.0.0.1:8080/execute",
                    {'helper':this.selected,
                     'supportitem':this.selected_support,
                     'gokigen':this.selected_gokigen
                    },{headers:headers
                    })
                    .then(function(res) {
                    storage.setItem('point_data',JSON.stringify(res.data));
                    });
                    this.server_flg = 1;
                    console.log("server : ok")
                }
                getSessionStorageData = JSON.parse(storage.getItem('point_data'));
                console.log(getSessionStorageData);
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
                    this.isnegative();
                }
            },
            reset() {
                this.server_flg = 0;
                this.selected_num_of_tornament = '';
                this.selected_gokigen = '';
                this.selected_character = '';
                this.selected_support = '';
                this.count_hit = 0;
                this.count_twobase = 0;
                this.count_threebase = 0;
                this.count_homerun = 0;
                this.count_bunt = 0;
                this.count_fly = 0;
                this.count_steal = 0;
                this.count_inings = 0;
                this.count_sunsinwithfastball = 0;
                this.count_gorowithfastball = 0;
                this.count_flywithfastball = 0;
                this.count_sunsinwithbreakingball = 0;
                this.count_gorowithbreakingball = 0;
                this.count_flywithbreakingball = 0;
                this.count_gettsu = 0;
                this.goukei = {'筋力':0,'敏捷':0,'技術':0,'変化球':0,'精神':0};
                storage.clear();
            },
            isnegative() {
                if (this.goukei['筋力'] <= 0){
                    this.goukei['筋力'] = 0;
                }
                if (this.goukei['敏捷'] <= 0){
                    this.goukei['敏捷'] = 0;
                }
                if (this.goukei['技術'] <= 0){
                    this.goukei['技術'] = 0;
                }
                if (this.goukei['変化球'] <= 0){
                    this.goukei['変化球'] = 0;
                }
                if (this.goukei['精神'] <= 0){
                    this.goukei['精神'] = 0;
                }
            }
          }
    }
</script>

<style lang="scss" scoped>
::v-deep .v-select__selections {
  input {
    width: 0;
  }
}

.balloon6 {
    position: relative;
    top: 50px;
    left: 150px;
    width: 100%;
    margin: 10px 0;
    overflow: hidden;
  }
  
  .balloon6 .faceicon {
    float: left;
    margin-right: -50px;
    width: 40px;
  }
  
  .balloon6 .faceicon img{
    width: 100%;
    height: auto;
    border-radius: 50%;
  }
  .balloon6 .chatting {
    width: 100%;
    text-align: left;
  }
  .says {
    display: inline-block;
    position: relative; 
    margin: 0 0 0 50px;
    padding: 10px;
    max-width: 250px;
    border-radius: 12px;
    background: #edf1ee;
  }
  
  .says:after {
    content: "";
    display: inline-block;
    position: absolute;
    top: 3px; 
    left: -19px;
    border: 8px solid transparent;
    border-right: 18px solid #edf1ee;
    -webkit-transform: rotate(35deg);
    transform: rotate(35deg);
  }
  .says p {
    margin: 0;
    padding: 0;
  }
  .v-select {
      font-size: 1.0em;
  }
</style>