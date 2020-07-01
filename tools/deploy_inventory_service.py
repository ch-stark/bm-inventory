import os
import utils
import argparse
import yaml

parser = argparse.ArgumentParser()
parser.add_argument("--target")
parser.add_argument("--domain")
parser.add_argument('--service', help='Service name to use', type=str, default=None)
parser.add_argument('--namespace', help='Namespace to use', type=str, default=None)
args = parser.parse_args()

SERVICE_NAME = args.name or 'assisted-installer'
NAMESPACE = args.namespace or 'assisted-installer'


def main():
    src_file = os.path.join(os.getcwd(), "deploy/bm-inventory-service.yaml")
    dst_file = os.path.join(os.getcwd(), "build/bm-inventory-service.yaml")
    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = yaml.safe_load(src)
            utils.update_metadata(data, name=args.service, namespace=args.namespace)
            print("Deploying {}".format(dst_file))
            yaml.dump(data, dst, default_flow_style=False)

    utils.apply(dst_file)

    # in case of openshift deploy ingress as well
    if args.target != "oc-ingress":
        src_file = os.path.join(os.getcwd(), "deploy/assisted-installer-ingress.yaml")
        dst_file = os.path.join(os.getcwd(), "build/assisted-installer-ingress.yaml")
        with open(src_file, "r") as src:
            with open(dst_file, "w+") as dst:
                data = src.read()
                data = data.replace("REPLACE_HOSTNAME",
                                    utils.get_service_host(SERVICE_NAME, args.target, args.domain, NAMESPACE))
                data = yaml.safe_load(data)
                utils.update_metadata(data, name=args.service, namespace=args.namespace)
                print("Deploying {}".format(dst_file))
                yaml.dump(data, dst, default_flow_style=False)
        utils.apply(dst_file)


if __name__ == "__main__":
    main()
