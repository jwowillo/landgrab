import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

@Component(
  selector: 'api',
  templateUrl: 'template.html',
  directives: const [ResourceComponent, ControllerDescriptionComponent],
  pipes: const [PrettyJSONPipe],
  providers: const [SchemaService],
)
class APIComponent implements OnInit {
  final List<Resource> resources = [];
  final List<Action> actions = [];

  String status;

  SchemaService _service;

  APIComponent(this._service);

  @override
  Future ngOnInit() async {
    try {
      resources.addAll(await _service.resources());
      actions.addAll(await _service.actions());
    } catch (error) {
      status = error.toString();
      print(error);
    }
  }
}
