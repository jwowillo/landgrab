import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/model/rules.dart';
import 'package:landgrab/service/api.dart';
import 'package:landgrab/service/convert.dart';

@Injectable()
class RulesService {
  Future<Rules> rules() async {
    Map<String, dynamic> json = await api('/rules');
    return mapToRules(json['data']['rules']);
  }
}
