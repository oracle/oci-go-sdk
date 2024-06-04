# Add or update spec utilities

import collections
import re
import xml.etree.ElementTree as ET


ns = {"ns": "http://maven.apache.org/POM/4.0.0"}


class CommentedTreeBuilder(ET.TreeBuilder):
    def __init__(self, *args, **kwargs):
        super(CommentedTreeBuilder, self).__init__(*args, **kwargs)

    def comment(self, data):
        self.start(ET.Comment, {})
        self.data(data)
        self.end(ET.Comment)


AddOrUpdateSpecResult = collections.namedtuple(
    'AddOrUpdateSpecResult',
    'updated existing ignored previous changed')


def compute_changed_settings(previous, current):
    changed = []
    for key, value in previous.items():
        if key not in current:
            changed.append(key)
        elif previous[key] != current[key]:
            changed.append(key)

    for key, value in current.items():
        if key not in previous:
            changed.append(key)

    return changed


def indent(elem, level=0):
    indent_str = "  "
    i = "\n" + level * indent_str
    if len(elem):
        if not elem.text or not elem.text.strip():
            elem.text = i + indent_str
        for e in elem:
            indent(e, level + 1)
            if not e.tail or not e.tail.strip():
                e.tail = i + indent_str
        if not e.tail or not e.tail.strip():
            e.tail = i
    else:
        if level and (not elem.tail or not elem.tail.strip()):
            elem.tail = i


def add_spec_module_to_github_whitelist(spec_name, github_whitelist_location, github_whitelist_template):
    if github_whitelist_location and github_whitelist_template:
        with open(github_whitelist_location, 'a') as f:
            f.write(github_whitelist_template.format(spec_name))


def convert_camel_to_snake_case(name):
    s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', name)
    return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()


def parse_pom(file_name):
    return ET.parse(file_name, parser=ET.XMLParser(target=CommentedTreeBuilder()))


def write_xml(file_name, pom):
    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


def find_pom_version(pom_location):
    print("Parsing pom: {}".format(pom_location))
    pom = parse_pom(pom_location)
    return pom.findall(".//ns:version", ns)[0].text
